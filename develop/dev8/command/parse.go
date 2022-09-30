package command

import (
	"errors"
	"strings"
)

func ParseCommand(cmd string, input string) (string, error) {
	args := strings.Fields(cmd)

	if len(args) == 0 {
		return "", nil
	}

	if input != "" {
		args = append(args, input)
	}

	var res string
	var err error
	switch args[0] {
	case "cd":
		res, err = cdCommand(args[1:])
	case "echo":
		res, err = echoCommand(args[1:])
	case "ps":
		res, err = psCommand()
	case "pwd":
		res, err = pwdCommand()
	case "kill":
		res, err = killCommand(args[1:])
	case "fork":
		res, err = forkCommand(args[1:])
	case "exec":
		res, err = execCommand(args[1:])
	default:
		return "", errors.New("myshell: unknown command: " + args[0])
	}

	if err != nil {
		return "", err
	}

	return res, nil
}
