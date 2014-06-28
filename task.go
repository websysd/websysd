package main

import (
	"bytes"
	"fmt"
	"github.com/ian-kent/go-log/log"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var Workspaces = make(map[string]*Workspace)

type Workspace struct {
	Name        string
	Environment map[string]string
	Tasks       map[string]*Task
}

func NewWorkspace(name string, environment map[string]string) *Workspace {
	ws := &Workspace{
		Name:        name,
		Environment: environment,
		Tasks:       make(map[string]*Task),
	}
	if _, ok := ws.Environment["WORKSPACE"]; !ok {
		ws.Environment["WORKSPACE"] = name
	}
	Workspaces[name] = ws
	return ws
}

type Task struct {
	Id          int
	Name        string
	Command     string
	Executor    []string
	Environment map[string]string
	Stdout      string
	Stderr      string

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
	Stdout      string
	Stderr      string
	StdoutBuf   LogWriter
	StderrBuf   LogWriter
	Environment map[string]string
	Executor    []string
}

func (tr *TaskRun) String() string {
	return fmt.Sprintf("Pid %d", tr.Cmd.Process.Pid)
}

type LogWriter interface {
	Write(p []byte) (n int, err error)
	String() string
	Len() int64
	Close()
}

type FileLogWriter struct {
	filename string
	file     *os.File
}

func NewFileLogWriter(file string) (*FileLogWriter, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}

	flw := &FileLogWriter{
		filename: file,
		file:     f,
	}
	return flw, nil
}

func (flw FileLogWriter) Close() {
	flw.file.Close()
}

func (flw FileLogWriter) Write(p []byte) (n int, err error) {
	return flw.file.Write(p)
}

func (flw FileLogWriter) String() string {
	b, err := ioutil.ReadFile(flw.filename)
	if err == nil {
		return string(b)
	}
	return ""
}

func (flw FileLogWriter) Len() int64 {
	s, err := os.Stat(flw.filename)
	if err == nil {
		return s.Size()
	}
	return 0
}

type InMemoryLogWriter struct {
	buffer *bytes.Buffer
}

func NewInMemoryLogWriter() InMemoryLogWriter {
	imlw := InMemoryLogWriter{}
	imlw.buffer = new(bytes.Buffer)
	return imlw
}

func (imlw InMemoryLogWriter) Write(p []byte) (n int, err error) {
	return imlw.buffer.Write(p)
}

func (imlw InMemoryLogWriter) String() string {
	return imlw.buffer.String()
}

func (imlw InMemoryLogWriter) Len() int64 {
	return int64(imlw.buffer.Len())
}

func (imlw InMemoryLogWriter) Close() {

}

type Event struct {
	Time    time.Time
	Message string
}

func NewTask(workspace *Workspace, name string, executor []string, command string, environment map[string]string, service bool, stdout string, stderr string) *Task {
	environment = AddDefaultVars(environment)

	if _, ok := environment["TASK"]; !ok {
		environment["TASK"] = name
	}

	stdout = ReplaceVars(stdout, environment)
	stderr = ReplaceVars(stderr, environment)

	task := &Task{
		Name:        name,
		Command:     command,
		Environment: environment,
		TaskRuns:    make([]*TaskRun, 0),
		Service:     service,
		Executor:    executor,
		Stdout:      stdout,
		Stderr:      stderr,
	}

	if task.Service {
		task.Start()
	}

	workspace.Tasks[name] = task

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
	run := len(t.TaskRuns)

	c := t.Command
	c = ReplaceVars(c, t.Environment)

	var cmd *exec.Cmd
	if len(t.Executor) > 0 {
		cmd = exec.Command(t.Executor[0], append(t.Executor[1:], c)...)
	} else {
		bits := strings.Split(c, " ")
		cmd = exec.Command(bits[0], bits[1:]...)
	}

	vars := map[string]string{
		"TASK": strconv.Itoa(t.Id),
		"RUN":  strconv.Itoa(run),
	}
	stdout := ReplaceVars(t.Stdout, vars)
	stderr := ReplaceVars(t.Stderr, vars)

	tr := &TaskRun{
		Id:          run,
		Events:      make([]*Event, 0),
		Cmd:         cmd,
		Command:     t.Command,
		Environment: make(map[string]string),
		Stdout:      stdout,
		Stderr:      stderr,
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

	if len(tr.Stdout) > 0 {
		wr, err := NewFileLogWriter(tr.Stdout)
		if err != nil {
			log.Error("Unable to open file %s: %s", tr.Stdout, err.Error())
			tr.StdoutBuf = NewInMemoryLogWriter()
		} else {
			tr.StdoutBuf = wr
		}
	} else {
		tr.StdoutBuf = NewInMemoryLogWriter()
	}
	if len(tr.Stderr) > 0 {
		wr, err := NewFileLogWriter(tr.Stderr)
		if err != nil {
			log.Error("Unable to open file %s: %s", tr.Stderr, err.Error())
			tr.StderrBuf = NewInMemoryLogWriter()
		} else {
			tr.StderrBuf = wr
		}
	} else {
		tr.StderrBuf = NewInMemoryLogWriter()
	}

	for k, v := range tr.Environment {
		log.Info("Adding env var %s = %s", k, v)
		tr.Cmd.Env = append(tr.Cmd.Env, k+"="+v)
	}

	err = tr.Cmd.Start()
	if tr.Cmd.Process != nil {
		ev := &Event{time.Now(), fmt.Sprintf("Process %d started: %s", tr.Cmd.Process.Pid, tr.Command)}
		log.Info(ev.Message)
		tr.Events = append(tr.Events, ev)
	}
	if err != nil {
		tr.Error = err
		log.Error(err.Error())
		tr.StdoutBuf.Close()
		tr.StderrBuf.Close()
		return
	}
	go func() {
		go io.Copy(tr.StdoutBuf, stdout)
		go io.Copy(tr.StderrBuf, stderr)

		tr.Cmd.Wait()

		tr.StdoutBuf.Close()
		tr.StderrBuf.Close()

		log.Trace("STDOUT: %s", tr.StdoutBuf.String())
		log.Trace("STDERR: %s", tr.StderrBuf.String())

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
