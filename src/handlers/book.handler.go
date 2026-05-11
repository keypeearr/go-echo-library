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

func (bh *BookHandler) GetBookByID(c *echo.Context) error {
	id, err := utils.ParseUUID(c.Param("id"))
	if err != nil {
		return err
	}

	book, err := bh.bs.GetBookByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, book)
}

func (bh *BookHandler) CreateBook(c *echo.Context) error {
	var body types.CreateBookRequest
	if err := c.Bind(&body); err != nil {
		return err
	}

	id, err := bh.bs.CreateBook(c.Request().Context(), body)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusCreated, types.CreateBookResult{
		ID: id,
	})
}

func (bh *BookHandler) UpdateBook(c *echo.Context) error {
	id, err := utils.ParseUUID(c.Param("id"))
	if err != nil {
		return err
	}

	var body types.UpdateBookRequest
	if err := c.Bind(&body); err != nil {
		return err
	}

	book, err := bh.bs.UpdateBook(c.Request().Context(), id, body)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, book)
}

func (bh *BookHandler) DeleteBook(c *echo.Context) error {
	id, err := utils.ParseUUID(c.Param("id"))
	if err != nil {
		return err
	}

	if err := bh.bs.DeleteBook(c.Request().Context(), id); err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, map[string]any{
		"status": "success",
	})
}
