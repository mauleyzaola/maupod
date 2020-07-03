package helpers

import (
	"crypto/sha512"
	"fmt"
	"io"
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
