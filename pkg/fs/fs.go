package fs

import (
	"github.com/nomionz/fs-enc/pkg/cipher"
	"os"
)

func Decrypt(path string, pass []byte) error {
	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	dec, err := cipher.Decrypt(file, pass)

	if err != nil {
		return err
	}

	return os.WriteFile(path, dec, 0644)
}

func Encrypt(path string, pass []byte) error {
	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	enc, err := cipher.Encrypt(file, pass)

	if err != nil {
		return err
	}

	return os.WriteFile(path, enc, 0644)
}
