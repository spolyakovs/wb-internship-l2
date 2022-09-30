package command

import (
	"fmt"
	"os"
)

func pwdCommand() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("pwd: %w", err)
	}
	return pwd, nil
}
