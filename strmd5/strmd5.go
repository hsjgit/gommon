package strmd5

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func StrMd5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func StrSha256(s string) string {
	shaCtx := sha256.New()
	shaCtx.Write([]byte(s))
	return hex.EncodeToString(shaCtx.Sum(nil))
}
