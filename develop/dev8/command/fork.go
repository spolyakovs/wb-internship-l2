package command

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func forkCommand(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("fork: enter process to fork")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", errors.New("fork: couldn't get pwd")
	}

	cmd, err := exec.LookPath(args[0])
	if err != nil {
		return "", fmt.Errorf("fork: couldn't find path for %v: %w", args[0], err)
	}
	if cmd == "" {
		return "", fmt.Errorf("fork: couldn't find path for %v", args[0])
	}
	args[0] = cmd

	_, err = syscall.ForkExec(args[0], args, &syscall.ProcAttr{
		Dir:   pwd,
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()}, // print message to the same pty
	})

	if err != nil {
		return "", fmt.Errorf("fork: could't fork: %w", err)
	}

	return "", nil
}
