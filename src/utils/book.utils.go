package utils

import (
	"time"
)

const (
	BookTitleMinLen = 1
	BookTitleMaxLen = 255

	BookIsbnMinLen = 10
	BookIsbnMaxLen = 13

	BookPublisherMinLen = 1
	BookPublisherMaxLen = 255

	BookLanguageMinLen = 1
	BookLanguageMaxLen = 100

	BookGenreMinLen = 1
	BookGenreMaxLen = 100

	BookDescriptionMinLen = 1
	BookDescriptionMaxLen = 5000

	BookPagesMin = int64(1)
	BookPagesMax = int64(100000)
)

func IsValidBookTitle(title string) bool {
	return IsLengthBetween(TrimString(title), BookTitleMinLen, BookTitleMaxLen)
}

func IsValidBookIsbn(isbn string) bool {
	return IsLengthBetween(TrimString(isbn), BookIsbnMinLen, BookIsbnMaxLen)
}

func IsValidBookPublisher(publisher string) bool {
	return IsLengthBetween(TrimString(publisher), BookPublisherMinLen, BookPublisherMaxLen)
}

func IsValidBookLanguage(language string) bool {
	return IsLengthBetween(TrimString(language), BookLanguageMinLen, BookLanguageMaxLen)
}

func IsValidBookGenre(genre string) bool {
	return IsLengthBetween(TrimString(genre), BookGenreMinLen, BookGenreMaxLen)
}

func IsValidBookDescription(description string) bool {
	return IsLengthBetween(TrimString(description), BookDescriptionMinLen, BookDescriptionMaxLen)
}

func IsValidBookPages(pages int64) bool {
	return pages >= BookPagesMin && pages <= BookPagesMax
}

func IsValidBookPublishedDate(t time.Time) bool {
	if t.IsZero() {
		return false
	}
	return !t.After(time.Now())
}
