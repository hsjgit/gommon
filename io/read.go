package io

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"sync/atomic"
	"time"

	"github.com/hsjgit/gommon/humanize"
)

// 一个具有计算hash功能的io.Reader
type HashReader struct {
	io.Reader
	hash    hash.Hash
	hashStr string
	fmtfun  func()
	ch      chan int64
	f       flowController
}

type flowController struct {
	speed int64
	limit int64
}

// 统计流量读取的速度
func (f *flowController) statistics(flow chan int64) {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println(fmt.Sprintf("当前速度为:%s/S", humanize.IBytes(uint64(atomic.LoadInt64(&f.speed)))))
			atomic.StoreInt64(&f.speed, 0)
		case num, ok := <-flow:
			if !ok {
				return
			}
			atomic.AddInt64(&f.speed, num)
		}
	}

}

func NewHashReader(r io.Reader, h hash.Hash) *HashReader {
	readhand := &HashReader{
		Reader: r,
		hash:   h,
		ch:     make(chan int64, 1),
		f:      flowController{},
	}
	go readhand.f.statistics(readhand.ch)
	return readhand
}

func (h *HashReader) Read(p []byte) (int, error) {
	if h.hash == nil {
		h.hash = md5.New()
	}

	n, err := h.Reader.Read(p)
	if n != 0 {
		io.Copy(h.hash, bytes.NewBuffer(p[:n]))
	}

	if h.fmtfun != nil {
		h.fmtfun()
	}
	if err != nil {
		h.hashStr = hex.EncodeToString(h.hash.Sum(nil))
		close(h.ch)
		return n, err
	}
	h.ch <- int64(n)
	return n, nil
}

func (h *HashReader) GetHashStr() string {
	return h.hashStr
}

func (h *HashReader) GetSeep() int64 {
	return atomic.LoadInt64(&h.f.speed)
}
