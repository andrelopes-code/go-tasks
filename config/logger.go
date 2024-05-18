package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug log.Logger
	info log.Logger
	warning log.Logger
	err log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)
	return &Logger{
		debug: *log.New(writer, "\033[1;42mDEBUG: \033[m ", logger.Flags()),
		info: *log.New(writer, "\033[1;46mINFO: \033[m ", logger.Flags()),
		warning: *log.New(writer, "\033[1;43mWARNING: \033[m ", logger.Flags()),
		err: *log.New(writer, "\033[1;41mERROR: \033[m ", logger.Flags()),
		writer: writer,
	}
}

// Create Non-Formatted logs
func (l *Logger) Debug(v ... interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ... interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warning(v ... interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ... interface{}) {
	l.err.Println(v...)
}
// Create Formatted logs
func (l *Logger) Debugf(format string, v ... interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ... interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warningf(format string, v ... interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ... interface{}) {
	l.err.Printf(format, v...)
}