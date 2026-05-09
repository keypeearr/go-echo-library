package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/kylerequez/go-echo-library/src/repositories"
	"github.com/kylerequez/go-echo-library/src/types"
	"github.com/kylerequez/go-echo-library/src/utils"
)

type AuthorService struct {
	queries *repositories.Queries
}

func NewAuthorService(q *repositories.Queries) *AuthorService {
	return &AuthorService{q}
}

func (as *AuthorService) GetAuthors(
	ctx context.Context,
	pagination types.BasePaginationRequest,
) (types.AuthorsPaginatedResult, error) {
	totalCount, err := as.queries.GetAuthorsCount(ctx)
	if err != nil {
		return types.AuthorsPaginatedResult{}, err
	}

	authors, err := as.queries.GetAuthors(ctx, repositories.GetAuthorsParams{
		Limit:  pagination.Limit,
		Offset: (pagination.Page - 1) * pagination.Limit,
	})
	if err != nil {
		return types.AuthorsPaginatedResult{}, err
	}

	var totalPages int64
	if totalCount%pagination.Limit != 0 {
		totalPages = (totalCount / pagination.Limit) + 1
	} else {
		totalPages = totalCount / pagination.Limit
	}

	return types.AuthorsPaginatedResult{
		Data: types.NewAuthorDTOs(authors),
		BasePaginatedResult: types.BasePaginatedResult{
			Count:      int64(len(authors)),
			Page:       pagination.Page,
			TotalCount: totalCount,
			TotalPages: totalPages,
		},
	}, nil
}

func (as *AuthorService) GetAuthorByID(
	ctx context.Context,
	id uuid.UUID,
) (types.AuthorDTO, error) {
	foundAuthor, err := as.queries.GetAuthorById(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return types.AuthorDTO{}, errors.New("author does not exist")
	}
	if err != nil {
		return types.AuthorDTO{}, err
	}

	return types.NewAuthorDTO(foundAuthor), nil
}

func (as *AuthorService) CreateAuthor(
	ctx context.Context,
	author types.CreateAuthorRequest,
) (uuid.UUID, error) {
	name := utils.TrimString(author.Name)
	biography := utils.TrimString(author.Biography)
	nationality := utils.TrimString(author.Nationality)
	image := utils.TrimString(*author.Image)

	if isValid := utils.IsValidAuthorBiography(biography); !isValid {
		return uuid.Nil, errors.New("biography is invalid")
	}

	if isValid := utils.IsValidAuthorNationality(nationality); !isValid {
		return uuid.Nil, errors.New("nationality is invalid")
	}

	if isValid := utils.IsValidAuthorBirthDate(author.BirthDate); !isValid {
		return uuid.Nil, errors.New("birthdate is invalid")
	}

	if isValid := utils.IsValidAuthorName(name); !isValid {
		return uuid.Nil, errors.New("name is invalid")
	}
	_, err := as.queries.GetAuthorByName(ctx, name)
	if err == nil {
		return uuid.Nil, errors.New("name already exists")
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return uuid.Nil, err
	}

	newAuthor := repositories.CreateAuthorParams{
		ID:          uuid.New(),
		Name:        name,
		Biography:   biography,
		BirthDate:   author.BirthDate,
		Nationality: nationality,
		Image:       sql.NullString{String: image, Valid: image != ""},
	}

	if err := as.queries.CreateAuthor(ctx, newAuthor); err != nil {
		return uuid.Nil, err
	}

	return newAuthor.ID, nil
}

func (as *AuthorService) UpdateAuthor(
	ctx context.Context,
	id uuid.UUID,
	author types.UpdateAuthorRequest,
) (types.AuthorDTO, error) {
	if author.Name == nil &&
		author.Biography == nil &&
		author.BirthDate == nil &&
		author.Nationality == nil &&
		author.Image == nil {
		return types.AuthorDTO{}, errors.New("no fields provided")
	}

	_, err := as.queries.GetAuthorById(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return types.AuthorDTO{}, errors.New("author does not exist")
	}
	if err != nil {
		return types.AuthorDTO{}, err
	}

	arg := repositories.UpdateAuthorParams{
		ID: id,
	}

	if author.Name != nil {
		name := utils.TrimString(*author.Name)

		if !utils.IsValidAuthorName(name) {
			return types.AuthorDTO{}, errors.New("name is invalid")
		}

		a, err := as.queries.GetAuthorByName(ctx, name)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return types.AuthorDTO{}, err
		}
		if err == nil && a.ID != id {
			return types.AuthorDTO{}, errors.New("name already exists")
		}

		arg.Name = sql.NullString{String: name, Valid: true}
	}

	if author.Biography != nil {
		biography := utils.TrimString(*author.Biography)

		if !utils.IsValidAuthorBiography(biography) {
			return types.AuthorDTO{}, errors.New("biography is invalid")
		}

		arg.Biography = sql.NullString{String: biography, Valid: true}
	}

	if author.Nationality != nil {
		nationality := utils.TrimString(*author.Nationality)

		if !utils.IsValidAuthorNationality(nationality) {
			return types.AuthorDTO{}, errors.New("nationality is invalid")
		}

		arg.Nationality = sql.NullString{String: nationality, Valid: true}
	}

	if author.BirthDate != nil {
		if !utils.IsValidAuthorBirthDate(*author.BirthDate) {
			return types.AuthorDTO{}, errors.New("birthdate is invalid")
		}

		arg.BirthDate = sql.NullTime{Time: *author.BirthDate, Valid: true}
	}

	if author.Image != nil {
		image := utils.TrimString(*author.Image)

		if !utils.IsValidAuthorImage(image) {
			return types.AuthorDTO{}, errors.New("image is invalid")
		}

		arg.Image = sql.NullString{String: image, Valid: true}
	}

	updatedAuthor, err := as.queries.UpdateAuthor(ctx, arg)
	if err != nil {
		return types.AuthorDTO{}, err
	}

	return types.NewAuthorDTO(updatedAuthor), nil
}

func (as *AuthorService) DeleteAuthor(ctx context.Context, id uuid.UUID) error {
	_, err := as.queries.GetAuthorById(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("author does not exist")
	}
	if err != nil {
		return err
	}

	if err := as.queries.DeleteAuthorById(ctx, id); err != nil {
		return err
	}

	return nil
}

func (as *AuthorService) GetAuthorBooks(
	ctx context.Context,
	id uuid.UUID,
	pagination types.BasePaginationRequest,
) (types.AuthorBooksPaginatedResult, error) {
	author, err := as.queries.GetAuthorById(ctx, id)
	if err != nil {
		return types.AuthorBooksPaginatedResult{}, err
	}

	totalCount, err := as.queries.GetBooksCountByAuthor(ctx, id)
	if err != nil {
		return types.AuthorBooksPaginatedResult{}, err
	}

	var totalPages int64
	if totalCount%pagination.Limit != 0 {
		totalPages = (totalCount / pagination.Limit) + 1
	} else {
		totalPages = totalCount / pagination.Limit
	}

	books, err := as.queries.GetBooksByAuthor(ctx, repositories.GetBooksByAuthorParams{
		AuthorID: id,
		Limit:    pagination.Limit,
		Offset:   (pagination.Page - 1) * pagination.Limit,
	})
	if err != nil {
		return types.AuthorBooksPaginatedResult{}, err
	}

	return types.AuthorBooksPaginatedResult{
		Author: types.NewAuthorDTO(author),
		Data:   types.NewBookDTOs(books),
		BasePaginatedResult: types.BasePaginatedResult{
			Count:      int64(len(books)),
			Page:       pagination.Page,
			TotalCount: totalCount,
			TotalPages: totalPages,
		},
	}, nil
}
