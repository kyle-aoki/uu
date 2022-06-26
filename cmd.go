package uu

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Cmd(command string) {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	commandParts := strings.Split(command, " ")
	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	cmd.Stdout = mw
	cmd.Stderr = mw
	err := cmd.Run()
	MustExec(err)
}
