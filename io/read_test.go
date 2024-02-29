package io

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
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

func EggFlower(src, dst string, container int64) ([]string, int64, error) {
	//if container < 8<<20 {
	//	container = 8 << 20
	//}
	address := make([]string, 0)
	fileStream, err := os.Open(src)
	if err != nil {
		return address, 0, err
	}
	if !strings.HasSuffix(dst, "/") {
		return address, 0, errors.New("dst 格式错误，应该以'/'结尾")
	}
	stat, err := fileStream.Stat()
	if err != nil {
		return nil, 0, err
	}
	i := 0
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return address, stat.Size(), err
	}
	count := int64(0)
	for {
		i++
		strBuff := strings.Builder{}
		strBuff.WriteString(dst)
		strBuff.WriteString(strconv.Itoa(i))
		strBuff.WriteString(".db")
		fileName := strBuff.String()
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			return address, stat.Size(), err
		}
		n, err := io.Copy(file, io.LimitReader(fileStream, container))
		if err != nil && err != io.EOF {
			return nil, 0, err
		}
		count += n
		address = append(address, fileName)
		if n <= container && count == stat.Size() {
			break
		}

	}
	return address, stat.Size(), nil
}

func TestMerge(t *testing.T) {
	dir, err := os.ReadDir("./hsj/")
	if err != nil {
		t.Log(err.Error())
		return
	}
	file, err := os.OpenFile("./hsj/merge.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		t.Log(err.Error())
		return
	}
	defer file.Close()
	fileName := make([]string, 0)
	for i := range dir {
		fileName = append(fileName, "./hsj/"+dir[i].Name())
	}
	sort.Strings(fileName)
	for i := range fileName {
		subfile, err := os.OpenFile(fileName[i], os.O_RDWR, os.ModePerm)
		if err != nil {
			t.Log(err.Error())
			return
		}
		io.Copy(file, subfile)
	}
}
