package io

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"testing"

	"github.com/hsjgit/gommon/gostring"
	"github.com/hsjgit/gommon/humanize"
)

func TestGetMd5(t *testing.T) {
	pipe, writer := io.Pipe()
	go writ(writer)

	r := NewHashReader(bufio.NewReader(pipe), sha256.New())
	b := make([]byte, 102400)
	for {
		_, err := r.Read(b)
		if err != nil {
			break
		}
		clear(b)
	}

}

func writ(writer io.Writer) {
	count := 0
	defer func() {
		fmt.Println(fmt.Sprintf("写入%d字节,%s", count, humanize.IBytes(uint64(count))))
	}()
	newWriter := bufio.NewWriter(writer)
	for {
		str := gostring.RandStr(102400, "")
		b := gostring.StringToBytes(str)
		newWriter.Write(b)
		count += len(b)
	}

}
