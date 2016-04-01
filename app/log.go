package app

import (
	"bytes"

	"github.com/ian-kent/go-log/appenders"
	"github.com/ian-kent/go-log/layout"
	"github.com/ian-kent/go-log/levels"
)

var maxlen = 262144000
var retain = 52428800

// Log is the in-memory application log for websysd
var Log bytes.Buffer

// Appender is a log appender
type Appender struct {
	a appenders.Appender
}

func (a *Appender) Write(level levels.LogLevel, message string, args ...interface{}) {
	a.a.Write(level, message, args...)
	Log.Write([]byte(a.Layout().Format(level, message, args...) + "\n"))
	if Log.Len() > maxlen {
		b := Log.Bytes()[retain:]
		Log = *new(bytes.Buffer)
		Log.Write(b)
	}
}

// SetLayout sets the layout
func (a *Appender) SetLayout(layout layout.Layout) {
	a.a.SetLayout(layout)
}

// Layout returns the layout
func (a *Appender) Layout() layout.Layout {
	return a.a.Layout()
}

// NewAppender returns a new appender
func NewAppender() *Appender {
	return &Appender{
		a: appenders.Console(),
	}
}
