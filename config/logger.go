package config

import (
	"io"
	"log"
	"os"
)

var (
	reset    = "\033[m"
	red      = "\033[1;31m"
	yellow   = "\033[1;33m"
	blue     = "\033[1;34m"
	green    = "\033[1;32m"
	bgred    = "\033[1;41m"
	bggreen  = "\033[1;42m"
	bgyellow = "\033[1;43m"
	bgblue   = "\033[1;44m"
)

type Logger struct {
	debug   log.Logger
	info    log.Logger
	warning log.Logger
	err     log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)
	return &Logger{
		debug:   *log.New(writer, bggreen+" [DEBUG] "+reset+green+" ", logger.Flags()),
		info:    *log.New(writer, bgblue+" [INFO] "+reset+blue+" ", logger.Flags()),
		warning: *log.New(writer, bgyellow+" [WARNING] "+reset+yellow+" ", logger.Flags()),
		err:     *log.New(writer, bgred+" [ERROR] "+reset+red+" ", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
	print(reset)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
	print(reset)
}
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
	print(reset)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
	print(reset)
}

// Create Formatted logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
	print(reset)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
	print(reset)
}
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
	print(reset)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
	print(reset)
}
