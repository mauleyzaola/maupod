package data

import (
	"fmt"
	"strings"
)

func LikeQuoted(value string) string {
	return fmt.Sprintf("%%%s%%", strings.ToLower(value))
}
