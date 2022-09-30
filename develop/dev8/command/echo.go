package command

import (
	"strings"
)

func echoCommand(args []string) (string, error) {
	return strings.Join(args, " "), nil
}
