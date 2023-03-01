package exec

import (
	"testing"
)

func TestShell(t *testing.T) {
	err := ShellAsyn("/Users/huangshijie/lotus/miner/start_miner.sh")
	if err != nil {
		t.Log(err.Error())
		return
	}

}
