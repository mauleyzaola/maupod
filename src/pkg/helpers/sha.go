package helpers

import (
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

func SHA(r io.Reader) ([]byte, error) {
	h := sha512.New512_256()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func HashFromSHA(data []byte) string {
	return fmt.Sprintf("%x", string(data))
}

func SHAFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()
	data, err := SHA(file)
	if err != nil {
		return "", err
	}
	return HashFromSHA(data), nil
}
