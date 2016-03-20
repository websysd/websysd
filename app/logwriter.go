package app

import (
	"bytes"
	"io/ioutil"
	"os"
)

// LogWriter is a log writer
type LogWriter interface {
	Write(p []byte) (n int, err error)
	String() string
	Len() int64
	Close()
}

// FileLogWriter is a log writer for files
type FileLogWriter struct {
	filename string
	file     *os.File
}

// NewFileLogWriter returns a new FileLogWriter
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

// Close closes the log writer
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

// Len returns the length of the file
func (flw FileLogWriter) Len() int64 {
	s, err := os.Stat(flw.filename)
	if err == nil {
		return s.Size()
	}
	return 0
}

// InMemoryLogWriter is an in memory log writer
type InMemoryLogWriter struct {
	buffer *bytes.Buffer
}

// NewInMemoryLogWriter returns a new InMemoryLogWriter
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

// Len returns the length of the content
func (imlw InMemoryLogWriter) Len() int64 {
	return int64(imlw.buffer.Len())
}

// Close closes the writer
func (imlw InMemoryLogWriter) Close() {

}
