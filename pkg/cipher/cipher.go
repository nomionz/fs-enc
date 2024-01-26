package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

const (
	keyLen     = 32
	saltSize   = 8
	iterations = 1000
)

func Encrypt(data, pass []byte) ([]byte, error) {
	salt := make([]byte, saltSize)

	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	gcm, err := newCipher(pass, salt)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encrypted := gcm.Seal(nonce, nonce, data, nil)
	encrypted = append(salt, encrypted...)

	return encrypted, nil
}

func Decrypt(data, pass []byte) ([]byte, error) {
	salt := data[:saltSize]
	data = data[saltSize:]

	gcm, err := newCipher(pass, salt)
	if err != nil {
		return nil, err
	}

	nsize := gcm.NonceSize()
	if len(data) < nsize {
		return nil, fmt.Errorf("data size and nonce size don't match")
	}

	nonce, txt := data[:nsize], data[nsize:]

	return gcm.Open(nil, nonce, txt, nil)
}

func deriveKey(pass, salt []byte) []byte {
	return pbkdf2.Key(pass, salt, iterations, keyLen, sha512.New)
}

func newCipher(pass, salt []byte) (cipher.AEAD, error) {
	key := deriveKey(pass, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(block)
}
