package main

import (
	"flag"
	"github.com/ian-kent/go-log/log"
	gotcha "github.com/ian-kent/gotcha/app"
	"github.com/ian-kent/gotcha/http"
	"net/url"
	"strconv"
)

func main() {
	global := "websysd.json"
	flag.StringVar(&global, "global", global, "global environment configuration")

	workspaces := make([]string, 0)
	flag.Var((*AppendSliceValue)(&workspaces), "workspace", "websysd workspace file (can be specified multiple times), defaults to './workspace.json'")

	// Create our Gotcha application
	var app = gotcha.Create(Asset)

	if len(workspaces) == 0 {
		workspaces = append(workspaces, "./workspace.json")
	}

	LoadConfig(global, workspaces)

	GlobalWorkspace = NewWorkspace(GlobalConfigWorkspace.Name, GlobalConfigWorkspace.Environment, make(map[string]map[string][]string))
	for fn, args := range GlobalConfigWorkspace.Functions {
		log.Info("=> Creating global function: %s", fn)
		GlobalWorkspace.Functions[fn] = &Function{
			Name:     fn,
			Args:     args.Args,
			Command:  args.Command,
			Executor: args.Executor,
		}
	}

	for _, ws := range ConfigWorkspaces {
		log.Info("=> Creating workspace: %s", ws.Name)

		var workspace *Workspace
		if wks, ok := Workspaces[ws.Name]; ok {
			log.Warn("Workspace %s already exists, merging tasks and environment")
			workspace = wks
		} else {
			workspace = NewWorkspace(ws.Name, ws.Environment, ws.Columns)
			Workspaces[ws.Name] = workspace
		}

		workspace.IsLocked = ws.IsLocked

		for fn, args := range ws.Functions {
			log.Info("=> Creating workspace function: %s", fn)
			workspace.Functions[fn] = &Function{
				Name:     fn,
				Args:     args.Args,
				Command:  args.Command,
				Executor: args.Executor,
			}
		}

		for _, t := range ws.Tasks {
			log.Info("=> Creating task: %s", t.Name)

			if _, ok := workspace.Tasks[t.Name]; ok {
				log.Warn("Task %s already exists, overwriting")
			}

			env := make(map[string]string)
			for k, v := range ws.Environment {
				env[k] = v
			}
			for k, v := range t.Environment {
				env[k] = v
			}

			task := NewTask(workspace, t.Name, t.Executor, t.Command, env, t.Service, t.Stdout, t.Stderr, t.Metadata)
			workspace.Tasks[t.Name] = task
		}
	}

	// Get the router
	r := app.Router

	// Create some routes
	r.Get("/", list_workspaces)
	r.Get("/workspace/(?P<workspace>[^/]+)", list_tasks)

	// Serve static content (but really use a CDN)
	r.Get("/images/(?P<file>.*)", r.Static("assets/images/{{file}}"))
	r.Get("/css/(?P<file>.*)", r.Static("assets/css/{{file}}"))

	r.Post("/workspace/(?P<workspace>[^/]+)/task/(?P<task>[^/]+)/start", startTask)
	r.Post("/workspace/(?P<workspace>[^/]+)/task/(?P<task>[^/]+)/stop", stopTask)
	r.Post("/workspace/(?P<workspace>[^/]+)/task/(?P<task>[^/]+)/enable", enableServiceTask)
	r.Post("/workspace/(?P<workspace>[^/]+)/task/(?P<task>[^/]+)/disable", disableServiceTask)
	r.Get("/workspace/(?P<workspace>[^/]+)/task/(?P<task>[^/]+)", taskHistory)

	r.Get("/workspace/(?P<workspace>[^/]+)/task/(?P<task>[^/]+)/run/(?P<run>\\d+)", taskRun)
	r.Get("/workspace/(?P<workspace>[^/]+)/task/(?P<task>[^/]+)/run/(?P<run>\\d+)/stdout", taskRunStdout)
	r.Get("/workspace/(?P<workspace>[^/]+)/task/(?P<task>[^/]+)/run/(?P<run>\\d+)/stderr", taskRunStderr)

	// Start our application
	app.Start()

	defer func() {
		for _, ws := range Workspaces {
			for _, t := range ws.Tasks {
				if t.ActiveTask != nil && t.ActiveTask.Cmd != nil && t.ActiveTask.Cmd.Process != nil {
					t.ActiveTask.Cmd.Process.Kill()
				}
			}
		}
	}()

	<-make(chan int)
}

func redir(session *http.Session) {
	redir := "/"

	if k := session.Request.Referer(); len(k) > 0 {
		redir = k
	}

	session.Redirect(&url.URL{Path: redir})
}

func startTask(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)
	id, _ := session.Stash["task"].(string)

	Workspaces[ws].Tasks[id].Start()

	redir(session)
}

func stopTask(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)
	id, _ := session.Stash["task"].(string)

	Workspaces[ws].Tasks[id].Stop()

	redir(session)
}

func enableServiceTask(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)
	id, _ := session.Stash["task"].(string)

	Workspaces[ws].Tasks[id].Service = true

	if Workspaces[ws].Tasks[id].ActiveTask == nil {
		Workspaces[ws].Tasks[id].Start()
	}

	redir(session)
}

func disableServiceTask(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)
	id, _ := session.Stash["task"].(string)

	Workspaces[ws].Tasks[id].Service = false

	redir(session)
}

func taskHistory(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)
	id, _ := session.Stash["task"].(string)

	session.Stash["Title"] = "Task"
	session.Stash["Page"] = "History"
	session.Stash["Workspace"] = Workspaces[ws]
	session.Stash["Task"] = Workspaces[ws].Tasks[id]

	session.RenderWithLayout("task.html", "layout.html", "Content")
}

func taskRun(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)
	id, _ := session.Stash["task"].(string)
	run, _ := strconv.Atoi(session.Stash["run"].(string))

	session.Stash["Title"] = "Task run"
	session.Stash["Page"] = "TaskRun"
	session.Stash["Workspace"] = Workspaces[ws]
	session.Stash["Task"] = Workspaces[ws].Tasks[id]
	session.Stash["TaskRun"] = Workspaces[ws].Tasks[id].TaskRuns[run]

	session.RenderWithLayout("taskrun.html", "layout.html", "Content")
}

func taskRunStdout(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)
	id, _ := session.Stash["task"].(string)
	run, _ := strconv.Atoi(session.Stash["run"].(string))

	session.Stash["Title"] = "Task run stdout"
	session.Stash["Page"] = "TaskOutput"
	session.Stash["Type"] = "stdout"
	session.Stash["Workspace"] = Workspaces[ws]
	session.Stash["Task"] = Workspaces[ws].Tasks[id]
	session.Stash["TaskRun"] = Workspaces[ws].Tasks[id].TaskRuns[run]
	session.Stash["LogOutput"] = Workspaces[ws].Tasks[id].TaskRuns[run].StdoutBuf.String()

	session.RenderWithLayout("log.html", "layout.html", "Content")
}

func taskRunStderr(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)
	id, _ := session.Stash["task"].(string)
	run, _ := strconv.Atoi(session.Stash["run"].(string))

	session.Stash["Title"] = "Task run stderr"
	session.Stash["Page"] = "TaskOutput"
	session.Stash["Type"] = "stderr"
	session.Stash["Workspace"] = Workspaces[ws]
	session.Stash["Task"] = Workspaces[ws].Tasks[id]
	session.Stash["TaskRun"] = Workspaces[ws].Tasks[id].TaskRuns[run]
	session.Stash["LogOutput"] = Workspaces[ws].Tasks[id].TaskRuns[run].StderrBuf.String()

	session.RenderWithLayout("log.html", "layout.html", "Content")
}

func list_workspaces(session *http.Session) {
	// Stash a value and render a template
	session.Stash["Title"] = "websysd"
	session.Stash["Page"] = "Workspaces"
	session.Stash["Workspaces"] = Workspaces
	session.RenderWithLayout("workspaces.html", "layout.html", "Content")
}

func list_tasks(session *http.Session) {
	ws, _ := session.Stash["workspace"].(string)

	// Stash a value and render a template
	session.Stash["Title"] = "websysd"
	session.Stash["Page"] = "Tasks"
	session.Stash["Workspace"] = Workspaces[ws]
	session.Stash["Tasks"] = Workspaces[ws].Tasks
	session.RenderWithLayout("tasks.html", "layout.html", "Content")
}
