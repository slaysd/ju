package util

import (
	"strings"
)

// ToString Join string array
func ToString(data []string) string {
	return strings.Join(data, " ")
}
