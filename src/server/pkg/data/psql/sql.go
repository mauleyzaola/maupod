package psql

import (
	"fmt"
	"strings"
)

func LikeQuoted(value string) string {
	return fmt.Sprintf("%%%s%%", strings.ToLower(value))
}
