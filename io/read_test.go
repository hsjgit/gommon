package io

import (
	"crypto/sha256"
	"io"
	"os"
	"testing"
)

func TestGetMd5(t *testing.T) {
	open, err := os.Open("/Users/huangshijie/test.sql")
	if err != nil {
		return
	}
	defer open.Close()
	r := NewHashReader(open, sha256.New())
	io.ReadAll(r)
	t.Log(r.GetHashStr())
}
