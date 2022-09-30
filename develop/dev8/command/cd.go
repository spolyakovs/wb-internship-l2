package command

import (
	"errors"
	"fmt"
	"os"
)

func cdCommand(args []string) (string, error) {
	switch len(args) {
	case 0:
		err := os.Chdir(os.Getenv("HOME"))
		if err != nil {
			return "", fmt.Errorf("cd: %w", err)
		}
	case 1:
		err := os.Chdir(args[0])
		if err != nil {
			return "", errors.New("cd: No such file or directory")
		}
	default:
		return "", errors.New("cd: too many arguments")
	}
	return "", nil
}
