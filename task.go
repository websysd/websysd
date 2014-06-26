package main

import (
	"bytes"
	"fmt"
	"github.com/ian-kent/go-log/log"
	"io"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type Task struct {
	Id          int
	Name        string
	Command     string
	Executor    []string
	Environment map[string]string

	ActiveTask *TaskRun
	TaskRuns   []*TaskRun

	Service bool
}

func (t *Task) GetExecutor() string {
	if len(t.Executor) == 0 {
		return "websysd"
	}
	return strings.Join(t.Executor, " ")
}

type TaskRun struct {
	Id          int
	Cmd         *exec.Cmd
	Error       error
	Started     time.Time
	Stopped     time.Time
	Events      []*Event
	Command     string
	StdoutBuf   bytes.Buffer
	StderrBuf   bytes.Buffer
	Environment map[string]string
	Executor    []string
}

func (tr *TaskRun) String() string {
	return fmt.Sprintf("Pid %d", tr.Cmd.Process.Pid)
}

type Event struct {
	Time    time.Time
	Message string
}

func NewTask(name string, executor []string, command string, environment map[string]string, service bool) *Task {
	task := &Task{
		Id:          len(Tasks),
		Name:        name,
		Command:     command,
		Environment: environment,
		TaskRuns:    make([]*TaskRun, 0),
		Service:     service,
		Executor:    executor,
	}
	Tasks = append(Tasks, task)
	TaskIndex[len(Tasks)-1] = task

	if task.Service {
		task.Start()
	}

	return task
}

func (t *Task) Start() {
	if t.ActiveTask == nil {
		t.ActiveTask = t.NewTaskRun()
		c := make(chan int)
		t.ActiveTask.Start(c)
		go func() {
			<-c
			t.ActiveTask = nil
			if t.Service {
				t.Start()
				return
			}
		}()
	}
}

func (t *Task) Stop() {
	if t.ActiveTask != nil {
		t.ActiveTask.Stop()
		t.ActiveTask = nil
	}
}

func (t *Task) NewTaskRun() *TaskRun {
	// FIXME needs improving (e.g. escaping - or maybe just use bash)
	c := t.Command
	for k, v := range t.Environment {
		c = strings.Replace(c, "$"+k, v, -1)
	}

	var cmd *exec.Cmd
	if len(t.Executor) > 0 {
		cmd = exec.Command(t.Executor[0], append(t.Executor[1:], c)...)
	} else {
		bits := strings.Split(c, " ")
		cmd = exec.Command(bits[0], bits[1:]...)
	}

	tr := &TaskRun{
		Id:          len(t.TaskRuns),
		Events:      make([]*Event, 0),
		Cmd:         cmd,
		Command:     t.Command,
		Environment: make(map[string]string),
	}

	for k, v := range t.Environment {
		tr.Environment[k] = v
	}

	t.TaskRuns = append(t.TaskRuns, tr)
	return tr
}

func (tr *TaskRun) Start(exitCh chan int) {
	tr.Started = time.Now()

	stdout, err := tr.Cmd.StdoutPipe()
	if err != nil {
		tr.Error = err
		return
	}
	stderr, err := tr.Cmd.StderrPipe()
	if err != nil {
		tr.Error = err
		return
	}

	for k, v := range tr.Environment {
		log.Info("Adding env var %s = %s", k, v)
		tr.Cmd.Env = append(tr.Cmd.Env, k + "=" + v)
	}

	err = tr.Cmd.Start()
	ev := &Event{time.Now(), fmt.Sprintf("Process %d started: %s", tr.Cmd.Process.Pid, tr.Command)}
	log.Info(ev.Message)
	tr.Events = append(tr.Events, ev)
	if err != nil {
		tr.Error = err
		log.Error(err.Error())
		return
	}
	go func() {
		go io.Copy(&tr.StdoutBuf, stdout)
		go io.Copy(&tr.StderrBuf, stderr)

		tr.Cmd.Wait()

		log.Info("STDOUT: %s", tr.StdoutBuf.String())
		log.Info("STDERR: %s", tr.StderrBuf.String())

		ps := tr.Cmd.ProcessState
		sy := ps.Sys().(syscall.WaitStatus)

		ev := &Event{time.Now(), fmt.Sprintf("Process %d exited with status %d", ps.Pid(), sy.ExitStatus())}
		log.Info(ev.Message)
		tr.Events = append(tr.Events, ev)
		log.Info(ps.String())

		tr.Stopped = time.Now()
		exitCh <- 1
	}()
}

func (tr *TaskRun) Stop() {
	if tr.Cmd == nil || tr.Cmd.Process == nil {
		return
	}

	tr.Cmd.Process.Kill()
}

func (t *Task) Status() string {
	if t.ActiveTask != nil && t.ActiveTask.Cmd != nil && t.ActiveTask.Cmd.Process != nil && t.ActiveTask.Cmd.Process.Pid > 0 {
		return "Running"
	}
	return "Stopped"
}

var Tasks = make([]*Task, 0)
var TaskIndex = make(map[int]*Task)
