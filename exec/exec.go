package exec

import (
	"io"
	"os/exec"
	"strings"
)

func Shell(shell string) (string, error) {
	command := exec.Command("/bin/bash", "-c", shell)
	pipe, err := command.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := command.Start(); err != nil {
		return "", err
	}

	all, err := io.ReadAll(pipe)
	if err != nil {
		return "", err
	}
	if err := command.Wait(); err != nil {
		return "", err
	}
	return string(all), nil
}

func ShellIO(shell string) (io.Reader, error) {
	command := exec.Command("/bin/bash", "-c", shell)
	pipe, err := command.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := command.Start(); err != nil {
		return nil, err
	}
	if err := command.Wait(); err != nil {
		return nil, err
	}
	return pipe, nil
}

func ShellLn(shell string) ([]string, error) {
	execShell, err := Shell(shell)
	if err != nil {
		return nil, err
	}
	split := strings.Split(execShell, "\n")
	return split, nil
}

// ShellAsyn 异步执行shell
func ShellAsyn(shell string) error {
	command := exec.Command("/bin/bash", "-c", shell)
	err := command.Start()
	if err != nil {
		return err
	}
	if err := command.Wait(); err != nil {
		return err
	}
	return nil
}
