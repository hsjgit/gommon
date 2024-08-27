package crypto

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"io"
	"os"
	"testing"

	goio "github.com/hsjgit/gommon/io"
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

func TestAESGCM(t *testing.T) {
	open, err := os.Open("image.jpg")
	if err != nil {
		t.Log(err.Error())
		return
	}
	defer open.Close()

	// 创建并打开输出文件
	file, err := os.OpenFile("image.gcm", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		t.Log(err.Error())
		return
	}
	defer file.Close()

	// 记录随机生成的 nonce
	t.Log(hex.EncodeToString(DefaultAESGCM.Nonce))
	reader := goio.NewHashReader(open, nil, nil)
	DefaultAESGCM.Reader = reader
	if _, err := io.Copy(file, DefaultAESGCM); err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(reader.GetHashStr()) // 67e7e0392a0816e256b974ad7af4a04c
	t.Log("File encryption successful")
}

func TestAESGCMDecryption(t *testing.T) {
	// 打开加密文件
	encryptedFile, err := os.Open("image.gcm")
	if err != nil {
		t.Fatalf("Failed to open encrypted file: %v", err)
	}
	defer func() {
		if cerr := encryptedFile.Close(); cerr != nil {
			t.Logf("Failed to close encrypted file: %v", cerr)
		}
	}()

	// 打开用于写入解密数据的文件
	decryptedFile, err := os.OpenFile("decrypted_image.jpg", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatalf("Failed to create decrypted file: %v", err)
	}
	defer func() {
		if cerr := decryptedFile.Close(); cerr != nil {
			t.Logf("Failed to close decrypted file: %v", cerr)
		}
	}()

	// 循环读取加密文件
	DefaultAESGCM.Nonce = make([]byte, 12)
	rand.Read(DefaultAESGCM.Nonce)
	reader := DefaultAESGCM.DecryptionReader(encryptedFile)
	for data := range reader {
		if data.Err != nil {
			t.Fatalf("Decryption error: %v", data.Err)
		}
		if data.Data == nil {
			break
		}
		if _, err := decryptedFile.Write(data.Data); err != nil {
			t.Fatalf("Failed to write decrypted data: %v", err)
		}
	}
	t.Log("File decryption successful")
}

// compareFiles 比较两个文件的内容是否相同
func compareFiles(file1, file2 *os.File) (bool, error) {
	buffer1 := make([]byte, 1024)
	buffer2 := make([]byte, 1024)

	for {
		n1, err1 := file1.Read(buffer1)
		n2, err2 := file2.Read(buffer2)

		if err1 != nil && err1 != io.EOF {
			return false, err1
		}
		if err2 != nil && err2 != io.EOF {
			return false, err2
		}

		if n1 != n2 || !bytes.Equal(buffer1[:n1], buffer2[:n2]) {
			return false, nil // 内容不同
		}

		// 文件读取完毕且内容相同
		if err1 == io.EOF && err2 == io.EOF {
			break
		}
	}
	return true, nil
}
