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
	configs := make([]string, 0)
	flag.Var((*AppendSliceValue)(&configs), "config", "websysd configuration file (can be specified multiple times), defaults to './websysd.json'")

	// Create our Gotcha application
	var app = gotcha.Create(Asset)

	if len(configs) == 0 {
		configs = append(configs, "./websysd.json")
	}

	// Load config
	for _, conf := range configs {
		log.Info("Loading configuration file: %s", conf)
		cfg, err := LoadConfig(conf)
		if err != nil {
			log.Error("Error loading config %s: %s", conf, err.Error())
		}
		if cfg != nil {
			for _, t := range cfg.Tasks {
				log.Info("=> Creating task: %s", t.Name)
				env := make(map[string]string)
				for k, v := range cfg.Environment {
					env[k] = v
				}
				for k, v := range t.Environment {
					env[k] = v
				}
				NewTask(t.Name, t.Executor, t.Command, env, t.Service, t.Stdout, t.Stderr)
			}
		}
	}

	// Get the router
	r := app.Router

	// Create some routes
	r.Get("/", tasks)

	// Serve static content (but really use a CDN)
	r.Get("/images/(?P<file>.*)", r.Static("assets/images/{{file}}"))
	r.Get("/css/(?P<file>.*)", r.Static("assets/css/{{file}}"))

	r.Post("/task/(?P<task>\\d+)/start", startTask)
	r.Post("/task/(?P<task>\\d+)/stop", stopTask)
	r.Post("/task/(?P<task>\\d+)/enable", enableServiceTask)
	r.Post("/task/(?P<task>\\d+)/disable", disableServiceTask)
	r.Get("/task/(?P<task>\\d+)", taskHistory)

	r.Get("/task/(?P<task>\\d+)/run/(?P<run>\\d+)", taskRun)
	r.Get("/task/(?P<task>\\d+)/run/(?P<run>\\d+)/stdout", taskRunStdout)
	r.Get("/task/(?P<task>\\d+)/run/(?P<run>\\d+)/stderr", taskRunStderr)

	// Start our application
	app.Start()

	defer func() {
		for _, t := range Tasks {
			if t.ActiveTask != nil && t.ActiveTask.Cmd != nil && t.ActiveTask.Cmd.Process != nil {
				t.ActiveTask.Cmd.Process.Kill()
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
	id, _ := strconv.Atoi(session.Stash["task"].(string))

	Tasks[id].Start()

	redir(session)
}

func stopTask(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))

	Tasks[id].Stop()

	redir(session)
}

func enableServiceTask(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))

	Tasks[id].Service = true

	if Tasks[id].ActiveTask == nil {
		Tasks[id].Start()
	}

	redir(session)
}

func disableServiceTask(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))

	Tasks[id].Service = false

	redir(session)
}

func taskHistory(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))

	session.Stash["Title"] = "Task"
	session.Stash["Page"] = "History"
	session.Stash["Task"] = Tasks[id]

	session.RenderWithLayout("task.html", "layout.html", "Content")
}

func taskRun(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))
	run, _ := strconv.Atoi(session.Stash["run"].(string))

	session.Stash["Title"] = "Task run"
	session.Stash["Page"] = "TaskRun"
	session.Stash["Task"] = Tasks[id]
	session.Stash["TaskRun"] = Tasks[id].TaskRuns[run]

	session.RenderWithLayout("taskrun.html", "layout.html", "Content")
}

func taskRunStdout(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))
	run, _ := strconv.Atoi(session.Stash["run"].(string))

	session.Stash["Title"] = "Task run stdout"
	session.Stash["Page"] = "TaskOutput"
	session.Stash["Type"] = "stdout"
	session.Stash["Task"] = Tasks[id]
	session.Stash["TaskRun"] = Tasks[id].TaskRuns[run]
	session.Stash["LogOutput"] = Tasks[id].TaskRuns[run].StdoutBuf.String()

	session.RenderWithLayout("log.html", "layout.html", "Content")
}

func taskRunStderr(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))
	run, _ := strconv.Atoi(session.Stash["run"].(string))

	session.Stash["Title"] = "Task run stderr"
	session.Stash["Page"] = "TaskOutput"
	session.Stash["Type"] = "stderr"
	session.Stash["Task"] = Tasks[id]
	session.Stash["TaskRun"] = Tasks[id].TaskRuns[run]
	session.Stash["LogOutput"] = Tasks[id].TaskRuns[run].StderrBuf.String()

	session.RenderWithLayout("log.html", "layout.html", "Content")
}

func tasks(session *http.Session) {
	// Stash a value and render a template
	session.Stash["Title"] = "websysd"
	session.Stash["Page"] = "Tasks"
	session.Stash["Tasks"] = Tasks
	session.RenderWithLayout("index.html", "layout.html", "Content")
}
