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

var GlobalWorkspace *Workspace
var Workspaces = make(map[string]*Workspace)

type Workspace struct {
	Name               string
	Environment        map[string]string
	Tasks              map[string]*Task
	IsLocked           bool
	Functions          map[string]*Function
	Columns            map[string]map[string][]string
	InheritEnvironment bool
}

func (ws *Workspace) GetColumn(task *Task, name string) string {
	log.Info("GetColumn: %s => %s", task.Name, name)
	col := ws.Columns[name]
	var fn []string
	var nm string
	for n, args := range col {
		nm = n
		fn = args
		break
	}
	return ws.ExecFunction(task, nm, fn...)
}

func (ws *Workspace) ExecFunction(task *Task, name string, args ...string) string {
	log.Info("Executing function %s: %s", name, args)
	var fn *Function
	if f, ok := ws.Functions[name]; ok {
		fn = f
	} else if f, ok := GlobalWorkspace.Functions[name]; ok {
		fn = f
	} else {
		log.Warn("Function not found: %s", name)
		return ""
	}

	argmap := make(map[string]string)
	for i, arg := range fn.Args {
		argmap[arg] = args[i]
	}

	for k, v := range argmap {
		log.Info("argmap: %s => %s", k, v)
		for t, m := range task.Metadata {
			log.Info("meta: %s => %s", t, m)
			v = strings.Replace(v, "$"+t, m, -1)
		}
		argmap[k] = v
	}

	c := fn.Command
	for k, v := range argmap {
		log.Info("ARG: %s => %s", k, v)
		c = strings.Replace(c, k, v, -1)
	}

	tsk := NewTask(nil, "Function$"+name, fn.Executor, c, make(map[string]string), false, "", "", make(map[string]string), "")
	ch := tsk.Start()
	<-ch
	return tsk.TaskRuns[0].StdoutBuf.String()
}

type Function struct {
	Name     string
	Args     []string
	Command  string
	Executor []string
}

func (ws *Workspace) ActiveTasks() int {
	a := 0
	for _, t := range ws.Tasks {
		if t.ActiveTask != nil {
			a++
		}
	}
	return a
}
func (ws *Workspace) InactiveTasks() int {
	return ws.TotalTasks() - ws.ActiveTasks()
}
func (ws *Workspace) TotalTasks() int {
	return len(ws.Tasks)
}
func (ws *Workspace) PercentActive() int {
	return int(float64(ws.ActiveTasks()) / float64(ws.TotalTasks()) * float64(100))
}
func (ws *Workspace) PercentInactive() int {
	return 100 - ws.PercentActive()
}

func NewWorkspace(name string, environment map[string]string, columns map[string]map[string][]string, inheritEnv bool) *Workspace {
	ws := &Workspace{
		Name:               name,
		Environment:        environment,
		Tasks:              make(map[string]*Task),
		Functions:          make(map[string]*Function),
		Columns:            columns,
		InheritEnvironment: inheritEnv,
	}
	if _, ok := ws.Environment["WORKSPACE"]; !ok {
		ws.Environment["WORKSPACE"] = name
	}
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
	Metadata    map[string]string
	Pwd         string

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
	WaitStatus  syscall.WaitStatus
	Pwd         string
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

func NewTask(workspace *Workspace, name string, executor []string, command string, environment map[string]string, service bool, stdout string, stderr string, metadata map[string]string, pwd string) *Task {
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
		Metadata:    metadata,
		Pwd:         pwd,
	}

	if task.Service {
		task.Start()
	}

	return task
}

func (t *Task) Start() chan int {
	c1 := make(chan int, 1)
	if t.ActiveTask == nil {
		t.ActiveTask = t.NewTaskRun()
		c := make(chan int)
		t.ActiveTask.Start(c)
		go func() {
			<-c
			c1 <- 1
			t.ActiveTask = nil
			if t.Service {
				t.Start()
				return
			}
		}()
	}
	return c1
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
		Pwd:         t.Pwd,
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
		exitCh <- 1
		return
	}
	stderr, err := tr.Cmd.StderrPipe()
	if err != nil {
		tr.Error = err
		exitCh <- 1
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

	if len(tr.Pwd) > 0 {
		log.Info("Setting pwd: %s", tr.Pwd)
		tr.Cmd.Dir = tr.Pwd
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
		exitCh <- 1
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
