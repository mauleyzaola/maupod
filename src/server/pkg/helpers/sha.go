package helpers

import (
	"crypto/sha512"
	"io"
)

func SHA(r io.Reader) ([]byte, error) {
	h := sha512.New512_256()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
