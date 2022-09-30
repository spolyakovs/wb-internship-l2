package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"github.com/spolyakovs/wb-internship-l2/develop/dev8/command"
)

type Shell struct {
	writer io.Writer
	reader io.Reader
}

func NewShell(writer io.Writer, reader io.Reader) *Shell {
	return &Shell{
		writer: writer,
		reader: reader,
	}
}

func (s *Shell) buildPrefix() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var prefix, postfix string
	user, err := user.Current()
	if err != nil {
		return "", errors.New("myshell: can not get current user info")
	}

	prefix = user.Username + `@myshell`

	home := os.Getenv("HOME")
	if strings.HasPrefix(path, home) {
		postfix = strings.Replace(path, home, "~", 1)
	} else {
		postfix = path
	}

	return greenBoldString(prefix) + ":" + blueBoldString(postfix) + "$ ", nil
}

func (s *Shell) Start() error {
	fmt.Fprint(s.writer, "entering myshell\n")
	fmt.Fprint(s.writer, `(type \exit to exit myshell)`+"\n")

	scanner := bufio.NewScanner(s.reader)

	var quit bool = false

	for {
		prefix, err := s.buildPrefix()
		if err != nil {
			return errors.New("myshell: can not check current directory place")
		}
		fmt.Fprint(s.writer, prefix)

		scanner.Scan()
		text := scanner.Text()

		cmds := strings.Split(text, " | ")
		_ = cmds

		var res string = "" // empty for the first cmd
		// err was declared before

		for _, cmd := range cmds {
			if cmd == `\exit` {
				quit = true
				break
			}

			// if multiple commands separated by "|", use previous result as last arg in this cmd
			res, err = command.ParseCommand(cmd, res)
			if err != nil {
				fmt.Fprintln(s.writer, redBoldString(err.Error()))
				break
			}
		}

		if res != "" {
			fmt.Fprintln(s.writer, res)
		}

		if quit {
			break
		}

		if scanner.Err() != nil {
			return errors.New("myshell: can not read data")
		}
	}

	if _, err := fmt.Fprintln(s.writer, "exiting myshell"); err != nil {
		return err
	}
	return nil
}
