package app

import (
	"strings"

	"github.com/ian-kent/go-log/log"
)

// GlobalWorkspace is the global workspace
var GlobalWorkspace *Workspace

// Workspaces is the user workspaces
var Workspaces = make(map[string]*Workspace)

// Workspace is a user workspace
type Workspace struct {
	Name               string
	Environment        map[string]string
	Tasks              map[string]*Task
	IsLocked           bool
	Functions          map[string]*Function
	Columns            map[string]map[string][]string
	InheritEnvironment bool
}

// GetColumn returns a workspace column
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

// ExecFunction executes a function task
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

	var funcEnvironment map[string]string
	if ws.InheritEnvironment {
		funcEnvironment = ws.Environment
	} else if GlobalWorkspace.InheritEnvironment {
		funcEnvironment = GlobalWorkspace.Environment
	} else {
		funcEnvironment = make(map[string]string)
	}

	tsk := NewTask(nil, "Function$"+name, fn.Executor, c, funcEnvironment, false, "", "", make(map[string]string), "")
	ch := tsk.Start()
	<-ch
	return tsk.TaskRuns[0].StdoutBuf.String()
}

// ActiveTasks returns the number of active tasks in a workspace
func (ws *Workspace) ActiveTasks() int {
	a := 0
	for _, t := range ws.Tasks {
		if t.ActiveTask != nil {
			a++
		}
	}
	return a
}

// InactiveTasks returns the number of inactive tasks in a workspace
func (ws *Workspace) InactiveTasks() int {
	return ws.TotalTasks() - ws.ActiveTasks()
}

// TotalTasks returns the total number of tasks in a workspace
func (ws *Workspace) TotalTasks() int {
	return len(ws.Tasks)
}

// PercentActive returns the percentage of tasks active in a workspace
func (ws *Workspace) PercentActive() int {
	return int(float64(ws.ActiveTasks()) / float64(ws.TotalTasks()) * float64(100))
}

// PercentInactive returns the percentage of tasks inactive in a workspace
func (ws *Workspace) PercentInactive() int {
	return 100 - ws.PercentActive()
}

// NewWorkspace returns a new workspace
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
