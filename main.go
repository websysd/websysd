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
				NewTask(t.Name, t.Command, t.Environment, t.Service, t.Enabled)
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
	r.Get("/task/(?P<task>\\d+)/history", taskHistory)

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

func startTask(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))

	Tasks[id].Start()

	session.Redirect(&url.URL{Path: "/"})
}

func stopTask(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))

	Tasks[id].Stop()

	session.Redirect(&url.URL{Path: "/"})
}

func taskHistory(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))

	session.Stash["Title"] = "Task history"
	session.Stash["Page"] = "History"
	session.Stash["Task"] = Tasks[id]

	session.RenderWithLayout("history.html", "layout.html", "Content")
}

func taskRunStdout(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))
	run, _ := strconv.Atoi(session.Stash["run"].(string))

	session.Stash["Title"] = "Task run stdout"
	session.Stash["Page"] = "History"
	session.Stash["Type"] = "stdout"
	session.Stash["LogOutput"] = Tasks[id].TaskRuns[run].StdoutBuf.String()

	session.RenderWithLayout("log.html", "layout.html", "Content")
}

func taskRunStderr(session *http.Session) {
	id, _ := strconv.Atoi(session.Stash["task"].(string))
	run, _ := strconv.Atoi(session.Stash["run"].(string))

	session.Stash["Title"] = "Task run stderr"
	session.Stash["Page"] = "History"
	session.Stash["Type"] = "stderr"
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
