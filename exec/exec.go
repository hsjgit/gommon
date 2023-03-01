package exec

import (
	"io/ioutil"
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

	all, err := ioutil.ReadAll(pipe)
	if err != nil {
		return "", err
	}
	if err := command.Wait(); err != nil {
		return "", err
	}
	return string(all), nil
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
	return nil
}
