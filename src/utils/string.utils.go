package utils

import (
	"strings"
	"unicode/utf8"
)

func TrimString(s string) string {
	return strings.TrimSpace(s)
}

func IsBlank(s string) bool {
	return TrimString(s) == ""
}

func IsLengthBetween(s string, min, max int) bool {
	n := utf8.RuneCountInString(s)
	return n >= min && n <= max
}
