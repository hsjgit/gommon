package exec

import (
	"io"
	"testing"
)

func TestShell(t *testing.T) {
	err := ShellAsyn("/Users/huangshijie/lotus/miner/start_miner.sh")
	if err != nil {
		t.Log(err.Error())
		return
	}

}

func TestShellIO(t *testing.T) {
	read, err := ShellIO("/Users/huangshijie/lotus/miner/start_miner.sh")
	if err != nil {
		t.Log(err.Error())
	}
	io.ReadAll(read)

}
