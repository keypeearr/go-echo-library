package utils

import (
	"strconv"

	"github.com/kylerequez/go-echo-library/src/types"
)

const (
	DefaultPaginationPage  = 1
	DefaultPaginationLimit = 20
	MaxPaginationLimit     = 100
)

func ResolvePagination(pageStr, limitStr string) types.BasePaginationRequest {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = DefaultPaginationPage
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = DefaultPaginationLimit
	}
	if limit > MaxPaginationLimit {
		limit = MaxPaginationLimit
	}

	return types.BasePaginationRequest{
		Page:  int64(page),
		Limit: int64(limit),
	}
}
