package fs

import (
	"github.com/nomionz/fs-enc/pkg/cipher"
	"os"
)

type cipherFunc func(data, pass []byte) ([]byte, error)

func Decrypt(path string, pass []byte) error {
	return processFile(path, pass, cipher.Decrypt)
}

func Encrypt(path string, pass []byte) error {
	return processFile(path, pass, cipher.Encrypt)
}

func processFile(path string, pass []byte, process cipherFunc) error {
	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	result, err := process(file, pass)

	if err != nil {
		return err
	}

	return os.WriteFile(path, result, 0644)
}
