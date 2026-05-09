package utils

import "strings"

func NormalizeISBN(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == 'x' || r == 'X':
			b.WriteRune('X')
		}
	}
	return b.String()
}

func IsValidISBN10(s string) bool {
	if len(s) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < 9; i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
		sum += int(c-'0') * (10 - i)
	}
	last := s[9]
	switch {
	case last == 'X':
		sum += 10
	case last >= '0' && last <= '9':
		sum += int(last - '0')
	default:
		return false
	}
	return sum%11 == 0
}

func IsValidISBN13(s string) bool {
	if len(s) != 13 {
		return false
	}
	sum := 0
	for i := 0; i < 13; i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if i%2 == 0 {
			sum += d
		} else {
			sum += d * 3
		}
	}
	return sum%10 == 0
}

func IsValidISBN(s string) bool {
	n := NormalizeISBN(s)
	return IsValidISBN10(n) || IsValidISBN13(n)
}
