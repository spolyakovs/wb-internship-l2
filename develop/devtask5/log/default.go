package logger

import (
	"fmt"
)

// TODO: comments
type loggerDefault struct {
	filename string
	needNum  bool
	output   []string
}

func NewLoggerDefault(filename string, needNum bool) *loggerDefault {
	return &loggerDefault{
		filename: filename,
		output:   make([]string, 0),
		needNum:  needNum,
	}
}

func (l *loggerDefault) Put(num int, line string, matched bool) {
	var separator string
	if matched {
		separator = ":"
	} else {
		separator = "-"
	}
	separator = cyanColor + separator + resetColor

	var prefix string = ""

	// num = 0 if line == "--"
	if num == 0 {
		line = fmt.Sprintf("%s%s%s", cyanColor, line, resetColor)
		l.output = append(l.output, line)
		return
	}

	if l.needNum {
		prefix = fmt.Sprintf("%s%v%s%s", greenColor, num, resetColor, separator)
	}
	if l.filename != "" && num != 0 {
		prefix = fmt.Sprintf("%s%v%s%s%s", purpleColor, l.filename, resetColor, separator, prefix)
	}

	if prefix != "" {
		line = fmt.Sprint(prefix + line)
	}

	l.output = append(l.output, line)
}

func (l *loggerDefault) Log() {
	for _, line := range l.output {
		fmt.Println(line)
	}
}
