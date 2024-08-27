package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

type AESGCM struct {
	io.Reader
	Buffer int64
	Key    []byte
	Nonce  []byte
	AEAD   cipher.AEAD
}

// GCMDecryptionData
// 解密的结构体
type GCMDecryptionData struct {
	Data []byte
	Err  error
}

// AES-GCM should be used because the operation is an authenticated encryption
// algorithm designed to provide both data authenticity (integrity) as well as
// confidentiality.

// Merged into Golang in https://go-review.googlesource.com/#/c/18803/

var defaultAESGCMKey = []byte("Sghi%^$%gh342uid") // 16 bytes for AES-128
var defaultAESGCMNonce = make([]byte, 12)         // 12 bytes nonce for GCM

// Generate a secure random nonce for the default AESGCM instance
func init() {
	if _, err := io.ReadFull(rand.Reader, defaultAESGCMNonce); err != nil {
		panic(fmt.Sprintf("failed to generate nonce: %v", err))
	}
}

// DefaultAESGCM
// 默认key长度为128位 16字节
// 默认nonce长度为96位 12字节
var DefaultAESGCM = &AESGCM{
	Key:    defaultAESGCMKey,
	Nonce:  defaultAESGCMNonce,
	Buffer: 1024,
}

func NewAESGCM(key, nonce []byte, read io.Reader, buffer int64) (*AESGCM, error) {
	aesgcm := &AESGCM{
		Reader: read,
		Key:    key,
		Buffer: buffer,
		Nonce:  nonce,
	}
	if err := aesgcm.init(); err != nil {
		return nil, err
	}
	return aesgcm, nil
}

func (a *AESGCM) init() error {
	if len(a.Key) != 16 && len(a.Key) != 32 {
		return errors.New("key size must be 16 or 32 bytes for AES-128 or AES-256")
	}
	if len(a.Nonce) != 12 {
		return errors.New("nonce size must be 12 bytes for GCM")
	}
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	a.AEAD = aesGCM
	return nil
}

func (a *AESGCM) GcmEncrypter(plaintext []byte) ([]byte, error) {
	if a.AEAD == nil {
		if err := a.init(); err != nil {
			return nil, err
		}
	}
	// Ensure the nonce is used only once with each key
	nonceCopy := make([]byte, len(a.Nonce))
	copy(nonceCopy, a.Nonce)
	ciphertext := a.AEAD.Seal(nil, nonceCopy, plaintext, nil)
	return ciphertext, nil
}

func (a *AESGCM) GCMDecrypter(ciphertext []byte) ([]byte, error) {
	if a.AEAD == nil {
		if err := a.init(); err != nil {
			return nil, err
		}
	}
	// Ensure the nonce is used correctly
	nonceCopy := make([]byte, len(a.Nonce))
	copy(nonceCopy, a.Nonce)
	plaintext, err := a.AEAD.Open(nil, nonceCopy, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}
	return plaintext, nil
}

func (a *AESGCM) Read(b []byte) (int, error) {
	if a.Reader == nil {
		return 0, errors.New("reader is nil")
	}

	// 暂时存放原始读取的数据
	if a.Buffer == 0 {
		a.Buffer = 1024
	}
	tmpBuffer := make([]byte, a.Buffer)
	n, err := a.Reader.Read(tmpBuffer)
	if err != nil && err != io.EOF {
		return 0, err
	}
	if n == 0 || err == io.EOF {
		return 0, io.EOF
	}

	// 生成新的 nonce
	nonceCopy := make([]byte, len(a.Nonce))
	copy(nonceCopy, a.Nonce)
	a.Nonce = nonceCopy
	encrypt, err := a.GcmEncrypter(tmpBuffer[:n])
	if err != nil {
		return 0, err
	}
	// 构建输出：nonce + 加密数据长度 + 加密数据
	length := len(encrypt)
	lengthBytes := make([]byte, 4)
	lengthBytes[0] = byte(length >> 24)
	lengthBytes[1] = byte(length >> 16)
	lengthBytes[2] = byte(length >> 8)
	lengthBytes[3] = byte(length)

	// 确保 b 的容量足够
	if len(b) < len(nonceCopy)+len(lengthBytes)+len(encrypt) {
		return 0, errors.New("buffer too small for encryption output")
	}

	n = 0
	n += copy(b[n:], nonceCopy)
	n += copy(b[n:], lengthBytes)
	n += copy(b[n:], encrypt)

	return n, nil
}

func (a *AESGCM) Close() error {
	if closer, ok := a.Reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// DecryptionReader
// 解密数据流
func (a *AESGCM) DecryptionReader(encryptedStream io.Reader) chan GCMDecryptionData {
	decryptedStream := make(chan GCMDecryptionData, 1)
	go a.decryptionReader(decryptedStream, encryptedStream)
	return decryptedStream
}

func (a *AESGCM) decryptionReader(decryptedStream chan GCMDecryptionData, encryptedStream io.Reader) {
	defer close(decryptedStream)
	for {
		nonce := make([]byte, 12)
		if _, err := io.ReadFull(encryptedStream, nonce); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			decryptedStream <- GCMDecryptionData{Err: fmt.Errorf("Failed to read nonce from file: %v", err)}
		}
		lengthBytes := make([]byte, 4)
		if _, err := io.ReadFull(encryptedStream, lengthBytes); err != nil {
			decryptedStream <- GCMDecryptionData{Err: fmt.Errorf("Failed to read block length: %v", err)}
			break
		}
		length := int(lengthBytes[0])<<24 | int(lengthBytes[1])<<16 | int(lengthBytes[2])<<8 | int(lengthBytes[3])

		encryptedData := make([]byte, length)
		if _, err := io.ReadFull(encryptedStream, encryptedData); err != nil {
			decryptedStream <- GCMDecryptionData{Err: fmt.Errorf("Failed to read encrypted data: %v", err)}
			break
		}
		a.Nonce = nonce
		decryptedData, err := a.GCMDecrypter(encryptedData)
		if err != nil {
			decryptedStream <- GCMDecryptionData{Err: fmt.Errorf("Decryption failed: %v", err)}
			break
		}
		decryptedStream <- GCMDecryptionData{Data: decryptedData}
	}
}
