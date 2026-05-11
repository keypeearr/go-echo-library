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

type BookService struct {
	queries *repositories.Queries
}

func NewBookService(q *repositories.Queries) *BookService {
	return &BookService{q}
}

func (bs *BookService) GetAllBooks(
	ctx context.Context,
	pagination types.BasePaginationRequest,
) (types.BooksPaginatedResult, error) {
	totalCount, err := bs.queries.GetBooksCount(ctx)
	if err != nil {
		return types.BooksPaginatedResult{}, err
	}

	books, err := bs.queries.GetBooks(ctx, repositories.GetBooksParams{
		Limit:  pagination.Limit,
		Offset: (pagination.Page - 1) * pagination.Limit,
	})
	if err != nil {
		return types.BooksPaginatedResult{}, err
	}

	var totalPages int64
	if totalCount%pagination.Limit != 0 {
		totalPages = (totalCount / pagination.Limit) + 1
	} else {
		totalPages = totalCount / pagination.Limit
	}

	return types.BooksPaginatedResult{
		Data: types.NewBookDTOs(books),
		BasePaginatedResult: types.BasePaginatedResult{
			Count:      int64(len(books)),
			Page:       pagination.Page,
			TotalCount: totalCount,
			TotalPages: totalPages,
		},
	}, nil
}

func (bs *BookService) GetBookByID(
	ctx context.Context,
	id uuid.UUID,
) (types.BookDTO, error) {
	book, err := bs.queries.GetBookById(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return types.BookDTO{}, errors.New("author does not exist")
	}
	if err != nil {
		return types.BookDTO{}, err
	}

	return types.NewBookDTO(book), nil
}

func (bs *BookService) CreateBook(
	ctx context.Context,
	book types.CreateBookRequest,
) (uuid.UUID, error) {
	title := utils.TrimString(book.Title)
	isbn := utils.TrimString(book.Isbn)
	publisher := utils.TrimString(*book.Publisher)
	publishedDate := book.PublishedDate
	pages := book.Pages
	language := utils.TrimString(*book.Language)
	genre := utils.TrimString(*book.Genre)
	description := utils.TrimString(*book.Description)

	if isValid := utils.IsValidBookTitle(title); !isValid {
		return uuid.Nil, errors.New("title is invalid")
	}

	if isValid := utils.IsValidBookPublisher(publisher); !isValid {
		return uuid.Nil, errors.New("publisher is invalid")
	}

	if isValid := utils.IsValidBookPublishedDate(publishedDate); !isValid {
		return uuid.Nil, errors.New("published date is invalid")
	}

	if isValid := utils.IsValidBookPages(pages); !isValid {
		return uuid.Nil, errors.New("pages is invalid")
	}

	if isValid := utils.IsValidBookLanguage(language); !isValid {
		return uuid.Nil, errors.New("language is invalid")
	}

	if isValid := utils.IsValidBookGenre(genre); !isValid {
		return uuid.Nil, errors.New("genre is invalid")
	}

	if isValid := utils.IsValidBookDescription(description); !isValid {
		return uuid.Nil, errors.New("description is invalid")
	}

	if isValid := utils.IsValidBookIsbn(isbn); !isValid {
		return uuid.Nil, errors.New("isbn is invalid")
	}

	_, err := bs.queries.GetBookByIsbn(ctx, isbn)
	if err == nil {
		return uuid.Nil, errors.New("isbn already exists")
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return uuid.Nil, err
	}

	arg := repositories.CreateBookParams{
		ID:            uuid.New(),
		Title:         title,
		AuthorID:      book.AuthorID,
		Isbn:          isbn,
		Publisher:     sql.NullString{String: publisher, Valid: publisher != ""},
		PublishedDate: publishedDate,
		Pages:         sql.NullInt64{Int64: pages, Valid: pages > 0},
		Language:      sql.NullString{String: language, Valid: language != ""},
		Genre:         sql.NullString{String: genre, Valid: genre != ""},
		Description:   sql.NullString{String: description, Valid: description != ""},
	}

	if err := bs.queries.CreateBook(ctx, arg); err != nil {
		return uuid.Nil, err
	}

	return arg.ID, nil
}

func (bs *BookService) UpdateBook(
	ctx context.Context,
	id uuid.UUID,
	body types.UpdateBookRequest,
) (types.BookDTO, error) {
	if body.Title == nil &&
		body.Isbn == nil &&
		body.Publisher == nil &&
		body.PublishedDate == nil &&
		body.Pages == nil &&
		body.Language == nil &&
		body.Genre == nil &&
		body.Description == nil {
		return types.BookDTO{}, errors.New("no fields provided")
	}

	arg := repositories.UpdateBookParams{
		ID: id,
	}

	if body.Title != nil {
		title := utils.TrimString(*body.Title)

		if isValid := utils.IsValidBookTitle(title); !isValid {
			return types.BookDTO{}, errors.New("title is invalid")
		}

		arg.Title = sql.NullString{String: title, Valid: title != ""}
	}

	if body.Publisher != nil {
		publisher := utils.TrimString(*body.Publisher)

		if isValid := utils.IsValidBookPublisher(publisher); !isValid {
			return types.BookDTO{}, errors.New("publisher is invalid")
		}

		arg.Publisher = sql.NullString{String: publisher, Valid: publisher != ""}
	}

	if body.PublishedDate != nil {
		publishedDate := body.PublishedDate

		if isValid := utils.IsValidBookPublishedDate(*publishedDate); !isValid {
			return types.BookDTO{}, errors.New("published date is invalid")
		}

		arg.PublishedDate = sql.NullTime{Time: *publishedDate, Valid: publishedDate != nil}
	}

	if body.Pages != nil {
		pages := body.Pages

		if isValid := utils.IsValidBookPages(*pages); !isValid {
			return types.BookDTO{}, errors.New("pages is invalid")
		}

		arg.Pages = sql.NullInt64{Int64: *pages, Valid: pages != nil}
	}

	if body.Language != nil {
		language := utils.TrimString(*body.Language)

		if isValid := utils.IsValidBookLanguage(language); !isValid {
			return types.BookDTO{}, errors.New("language is invalid")
		}

		arg.Language = sql.NullString{String: language, Valid: language != ""}
	}

	if body.Genre != nil {
		genre := utils.TrimString(*body.Genre)

		if isValid := utils.IsValidBookGenre(genre); !isValid {
			return types.BookDTO{}, errors.New("genre is invalid")
		}

		arg.Genre = sql.NullString{String: genre, Valid: genre != ""}
	}

	if body.Description != nil {
		description := utils.TrimString(*body.Description)

		if isValid := utils.IsValidBookDescription(description); !isValid {
			return types.BookDTO{}, errors.New("description is invalid")
		}

		arg.Description = sql.NullString{String: description, Valid: description != ""}
	}

	if body.Isbn != nil {
		isbn := utils.TrimString(*body.Isbn)

		if isValid := utils.IsValidBookIsbn(isbn); !isValid {
			return types.BookDTO{}, errors.New("isbn is invalid")
		}

		_, err := bs.queries.GetBookByIsbn(ctx, isbn)
		if err == nil {
			return types.BookDTO{}, errors.New("isbn already exists")
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return types.BookDTO{}, err
		}

		arg.Isbn = sql.NullString{String: isbn, Valid: isbn != ""}
	}

	book, err := bs.queries.UpdateBook(ctx, arg)
	if err != nil {
		return types.BookDTO{}, err
	}

	return types.NewBookDTO(book), nil
}

func (bs *BookService) DeleteBook(ctx context.Context, id uuid.UUID) error {
	_, err := bs.queries.GetBookById(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("book does not exist")
	}
	if err != nil {
		return err
	}

	if err := bs.queries.DeleteBook(ctx, id); err != nil {
		return err
	}

	return nil
}
