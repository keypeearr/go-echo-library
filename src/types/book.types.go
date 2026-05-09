package types

import (
	"time"

	"github.com/google/uuid"

	"github.com/kylerequez/go-echo-library/src/repositories"
)

type BookDTO struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	AuthorID      uuid.UUID `json:"author_id"`
	Isbn          string    `json:"isbn"`
	Publisher     *string   `json:"publisher"`
	PublishedDate time.Time `json:"published_date"`
	Pages         int64     `json:"pages"`
	Language      *string   `json:"language"`
	Genre         *string   `json:"genre"`
	Description   *string   `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type BooksPaginatedResult struct {
	Data []BookDTO `json:"data"`
	BasePaginatedResult
}

type CreateBookRequest struct {
	Title         string    `json:"title"`
	AuthorID      uuid.UUID `json:"author_id"`
	Isbn          string    `json:"isbn"`
	Publisher     *string   `json:"publisher"`
	PublishedDate time.Time `json:"published_date"`
	Pages         int64     `json:"pages"`
	Language      *string   `json:"language"`
	Genre         *string   `json:"genre"`
	Description   *string   `json:"description"`
}

type UpdateBookRequest struct{}

type CreateBookResult struct {
	ID uuid.UUID `json:"id"`
}

func NewBookDTO(book repositories.Book) BookDTO {
	var publisher *string
	if book.Publisher.Valid {
		s := book.Publisher.String
		publisher = &s
	}

	var language *string
	if book.Language.Valid {
		s := book.Language.String
		language = &s
	}

	var description *string
	if book.Description.Valid {
		s := book.Description.String
		description = &s
	}

	var genre *string
	if book.Genre.Valid {
		s := book.Genre.String
		genre = &s
	}

	return BookDTO{
		ID:            book.ID,
		Title:         book.Title,
		AuthorID:      book.AuthorID,
		Isbn:          book.Isbn,
		Publisher:     publisher,
		PublishedDate: book.PublishedDate,
		Pages:         book.Pages.Int64,
		Language:      language,
		Genre:         genre,
		Description:   description,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
	}
}

func NewBookDTOs(books []repositories.Book) []BookDTO {
	out := make([]BookDTO, len(books))
	for i, a := range books {
		out[i] = NewBookDTO(a)
	}
	return out
}
