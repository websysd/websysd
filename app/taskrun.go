package app

import (
	"fmt"
	"io"
	"os/exec"
	"syscall"
	"time"

	"github.com/ian-kent/go-log/log"
)

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

		ps := tr.Cmd.ProcessState
		sy := ps.Sys().(syscall.WaitStatus)

		if sy.ExitStatus() == 0 {
			log.Trace("STDOUT: %s", tr.StdoutBuf.String())
			log.Trace("STDERR: %s", tr.StderrBuf.String())
		} else {
			log.Info("STDOUT: %s", tr.StdoutBuf.String())
			log.Info("STDERR: %s", tr.StderrBuf.String())
		}

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
