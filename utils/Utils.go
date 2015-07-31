package utils

import (
	"strings"
)

func Concat(a string, b string) string {
	s := []string{a, b}
	return strings.Join(s, "")
}