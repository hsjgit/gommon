package crypto

import (
	"encoding/base32"
	"testing"
)

// 82b97f8395b54e4b1f1bdb9b53e1a2be
func TestAES(t *testing.T) {
	a, err := AesEncryptWithSalt([]byte("ajdJ:isJIAklsias"), []byte("huangshije"))
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(base32.StdEncoding.EncodeToString(a))
	decrypt, err := AesDecryptWithSalt([]byte("ajdJ:isJIAklsias"), a)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(string(decrypt))

}
