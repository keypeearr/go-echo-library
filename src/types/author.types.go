package types

import (
	"time"

	"github.com/google/uuid"

	"github.com/kylerequez/go-echo-library/src/repositories"
)

type AuthorDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Biography   string    `json:"biography"`
	BirthDate   time.Time `json:"birth_date"`
	Nationality string    `json:"nationality"`
	Image       *string   `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AuthorsPaginatedResult struct {
	Data []AuthorDTO `json:"data"`
	BasePaginatedResult
}

type AuthorBooksPaginatedResult struct {
	Author AuthorDTO `json:"author"`
	Data   []BookDTO `json:"data"`
	BasePaginatedResult
}

type CreateAuthorRequest struct {
	Name        string    `json:"name"`
	Biography   string    `json:"biography"`
	BirthDate   time.Time `json:"birth_date"`
	Nationality string    `json:"nationality"`
	Image       *string   `json:"image"`
}

type UpdateAuthorRequest struct {
	Name        *string    `json:"name,omitempty"`
	Biography   *string    `json:"biography,omitempty"`
	BirthDate   *time.Time `json:"birth_date,omitempty"`
	Nationality *string    `json:"nationality,omitempty"`
	Image       *string    `json:"image,omitempty"`
}

type CreateAuthorResult struct {
	ID uuid.UUID `json:"id"`
}

func NewAuthorDTO(author repositories.Author) AuthorDTO {
	var image *string
	if author.Image.Valid {
		s := author.Image.String
		image = &s
	}

	return AuthorDTO{
		ID:          author.ID,
		Name:        author.Name,
		Biography:   author.Biography,
		BirthDate:   author.BirthDate,
		Nationality: author.Nationality,
		Image:       image,
		CreatedAt:   author.CreatedAt,
		UpdatedAt:   author.UpdatedAt,
	}
}

func NewAuthorDTOs(authors []repositories.Author) []AuthorDTO {
	out := make([]AuthorDTO, len(authors))
	for i, a := range authors {
		out[i] = NewAuthorDTO(a)
	}
	return out
}
