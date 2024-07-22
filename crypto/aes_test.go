package crypto

import (
	"encoding/base32"
	"testing"
)

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
