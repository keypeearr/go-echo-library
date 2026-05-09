package services

import (
	"context"

	"github.com/kylerequez/go-echo-library/src/repositories"
	"github.com/kylerequez/go-echo-library/src/types"
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
