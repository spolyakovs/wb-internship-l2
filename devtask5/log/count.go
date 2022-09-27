package logger

import (
	"fmt"
	"strconv"
)

// TODO: comments
type loggerCount struct {
	filename string
	counter  int
}

func NewLoggerCount(filename string) *loggerCount {
	return &loggerCount{
		filename: filename,
		counter:  0,
	}
}

func (l *loggerCount) Put(num int, line string, found bool) {
	if found {
		l.counter++
	}
}

func (l *loggerCount) Log() {
	str := strconv.Itoa(l.counter)
	var prefix string = ""
	if l.filename != "" {
		prefix = fmt.Sprintf("%s%s%s%s:%s", purpleColor, l.filename, resetColor, cyanColor, resetColor)
	}
	fmt.Println(prefix + str)
}
