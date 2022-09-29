package main

import "fmt"

type logLevel int

const (
	PanicLevel logLevel = iota
	WarnLevel
	InfoLevel
)

type Logger interface {
	Message(string, logLevel)
	SetNext(Logger)
}

type logger struct {
	next   Logger
	level  logLevel
	prefix string
}

func NewLogger(level logLevel, prefix string) *logger {
	return &logger{level: level, prefix: prefix}
}

func (l *logger) Message(msg string, level logLevel) {
	if level == l.level {
		fmt.Println(l.prefix, msg)
		return
	}
	if l.next != nil {
		l.next.Message(msg, level)
	}
}

func (l *logger) SetNext(next Logger) {
	l.next = next
}

func ChainExample() {
	p := NewLogger(PanicLevel, "PANIC")
	w := NewLogger(WarnLevel, "WARN")
	i := NewLogger(InfoLevel, "INFO")

	p.SetNext(w)
	w.SetNext(i)

	p.Message("panic message", PanicLevel)
	p.Message("info message", InfoLevel)
	p.Message("warning message", WarnLevel)
}
