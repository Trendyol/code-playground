package main

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	FATAL = "FATAL"
	ERROR = "ERROR"
	WARN  = "WARNING"
	INFO  = "INFO"
	DEBUG = "DEBUG"
	TRACE = "TRACE"
)

type ILoggers interface {
	Debug(message interface{})
	Info(message interface{})
	Warn(message interface{})
	Error(message interface{})
	Fatal(message interface{})
	Trace(message interface{})
	Log(level string, message interface{})
}

type Loggers struct{}

func NewLogger() ILoggers {
	return &Loggers{}
}

func (loggers *Loggers) Debug(message interface{}) {
	loggers.Log(DEBUG, message)
}

func (loggers *Loggers) Info(message interface{}) {
	loggers.Log(INFO, message)
}

func (loggers *Loggers) Warn(message interface{}) {
	loggers.Log(WARN, message)
}

func (loggers *Loggers) Error(message interface{}) {
	loggers.Log(ERROR, message)
}

func (loggers *Loggers) Fatal(message interface{}) {
	loggers.Log(FATAL, message)
}

func (loggers *Loggers) Trace(message interface{}) {
	loggers.Log(TRACE, message)
}

func (loggers *Loggers) Log(level string, message interface{}) {

	Spinner.Stop()

	c := color.New(color.FgWhite)

	if level == ERROR {
		c = color.New(color.FgRed, color.Bold)
	}

	if level == INFO {
		c = color.New(color.FgGreen, color.Bold)
	}

	_, _ = c.Println(fmt.Sprintf("⇨ %s", message))
}
