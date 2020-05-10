package helpers

import (
	"github.com/google/uuid"
)

func NewUUID() string {
	val, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return val.String()
}
