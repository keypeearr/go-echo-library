package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/kylerequez/go-echo-library/src/services"
	"github.com/kylerequez/go-echo-library/src/types"
	"github.com/kylerequez/go-echo-library/src/utils"
)

type BookHandler struct {
	bs *services.BookService
}

func NewBookHandler(bs *services.BookService) *BookHandler {
	return &BookHandler{bs}
}

func (bh *BookHandler) GetAllBooks(c *echo.Context) error {
	p := c.QueryParamOr("page", "")
	l := c.QueryParamOr("limit", "")

	pagination := utils.ResolvePagination(p, l)

	result, err := bh.bs.GetAllBooks(c.Request().Context(), pagination)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, result)
}

func (bh *BookHandler) CreateBook(c *echo.Context) error {
	var body types.CreateBookRequest
	if err := c.Bind(&body); err != nil {
		return err
	}

	return nil
}

func (bh *BookHandler) UpdateBook(c *echo.Context) error {
	return nil
}

func (bh *BookHandler) DeleteBook(c *echo.Context) error {
	return nil
}
