package logger

import "runtime"

var (
	resetColor  string
	greenColor  string
	cyanColor   string
	purpleColor string
)

func init() {
	if runtime.GOOS == "windows" {
		// windows console doesn't support colors
		resetColor = ""
		greenColor = ""
		cyanColor = ""
		purpleColor = ""
	} else {
		resetColor = "\033[0m"
		greenColor = "\033[32m"
		cyanColor = "\033[36m"
		purpleColor = "\033[35m"
	}
}

// TODO: comments
type Logger interface {
	Put(int, string, bool)
	Log()
}
