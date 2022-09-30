package main

import (
	"log"
	"os"
	"runtime"
)

var (
	redBoldColor   string // errors
	greenBoldColor string // prefix (username@myshell)
	blueBoldColor  string // wd
	resetColor     string
)

func init() {
	if runtime.GOOS == "windows" {
		// windows console doesn't support colors
		redBoldColor = ""
		greenBoldColor = ""
		blueBoldColor = ""
		resetColor = ""
	} else {
		redBoldColor = "\033[1m\033[31m"
		greenBoldColor = "\033[1m\033[32m"
		blueBoldColor = "\033[1m\033[34m"
		resetColor = "\033[0m"
	}
}

func main() {
	shell := NewShell(os.Stdout, os.Stdin)

	if err := shell.Start(); err != nil {
		log.Fatalln(err)
	}
}
