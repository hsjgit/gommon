package io

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"io"
)

// 一个具有计算hash功能的io.Reader
type HashReader struct {
	io.Reader
	hash    hash.Hash
	hashStr string
}

func NewHashReader(r io.Reader, h hash.Hash) *HashReader {
	return &HashReader{
		Reader: r,
		hash:   h,
	}
}

func (h *HashReader) Read(p []byte) (int, error) {
	if h.hash == nil {
		h.hash = md5.New()
	}
	n, err := h.Reader.Read(p)
	if n != 0 {
		io.Copy(h.hash, bytes.NewBuffer(p[:n]))
	}
	if err != nil {
		h.hashStr = hex.EncodeToString(h.hash.Sum(nil))
		return n, err
	}
	return n, nil
}

func (h *HashReader) GetHashStr() string {
	return h.hashStr
}
