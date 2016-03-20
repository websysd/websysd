package app

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Task represents an individual task
type Task struct {
	ID          int
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

// GetExecutor gets the task executor
func (t *Task) GetExecutor() string {
	if len(t.Executor) == 0 {
		return "websysd"
	}
	return strings.Join(t.Executor, " ")
}

// Event represents an event
type Event struct {
	Time    time.Time
	Message string
}

// NewTask returns a new Task
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

// Start starts a task
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
				time.Sleep(time.Second * 1)
				t.Start()
				return
			}
		}()
	}
	return c1
}

// Stop stops a task
func (t *Task) Stop() {
	if t.ActiveTask != nil {
		t.ActiveTask.Stop()
		t.ActiveTask = nil
	}
}

// NewTaskRun returns a new TaskRun for the Task
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
		"TASK": strconv.Itoa(t.ID),
		"RUN":  strconv.Itoa(run),
	}
	if len(t.Pwd) > 0 {
		vars["PWD"] = t.Pwd
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

// Status returns a string representation of the current task status
func (t *Task) Status() string {
	if t.ActiveTask != nil && t.ActiveTask.Cmd != nil && t.ActiveTask.Cmd.Process != nil && t.ActiveTask.Cmd.Process.Pid > 0 {
		return "Running"
	}
	return "Stopped"
}
