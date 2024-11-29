package utils

import (
	"fmt"
	"math/rand/v2"
	"os/exec"
)

var DataDir string

func GenerateID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := len(charset)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.IntN(charsetLength)]
	}
	return string(b)
}

func ExecCommand(command []string) ([]byte, error) {
	if len(command) == 0 {
		return nil, fmt.Errorf("no command provided")
	} else if len(command) == 1 {
		return exec.Command(command[0]).Output()
	}
	return exec.Command(command[0], command[1:]...).Output()
}
