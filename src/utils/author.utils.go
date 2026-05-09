package utils

import (
	"time"
)

const (
	AuthorNameMinLen = 1
	AuthorNameMaxLen = 255

	AuthorBiographyMinLen = 1
	AuthorBiographyMaxLen = 5000

	AuthorNationalityMinLen = 1
	AuthorNationalityMaxLen = 100

	AuthorImageMaxLen = 2048
)

func IsValidAuthorName(name string) bool {
	return IsLengthBetween(TrimString(name), AuthorNameMinLen, AuthorNameMaxLen)
}

func IsValidAuthorBiography(bio string) bool {
	return IsLengthBetween(TrimString(bio), AuthorBiographyMinLen, AuthorBiographyMaxLen)
}

func IsValidAuthorNationality(nat string) bool {
	return IsLengthBetween(TrimString(nat), AuthorNationalityMinLen, AuthorNationalityMaxLen)
}

func IsValidAuthorImage(s string) bool {
	return IsLengthBetween(TrimString(s), 0, AuthorImageMaxLen)
}

func IsValidAuthorBirthDate(t time.Time) bool {
	if t.IsZero() {
		return false
	}
	return !t.After(time.Now())
}
