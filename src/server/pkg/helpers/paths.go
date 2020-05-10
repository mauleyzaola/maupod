package helpers

import (
	"os"
	"path/filepath"
)

func PathBackend() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "mauleyzaola", "maupod", "src", "server")
}
