package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/kylerequez/go-echo-library/src/services"
	"github.com/kylerequez/go-echo-library/src/types"
	"github.com/kylerequez/go-echo-library/src/utils"
)

type AuthorHandler struct {
	as *services.AuthorService
}

func NewAuthorHandler(as *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{as}
}

func (ah *AuthorHandler) GetAuthors(c *echo.Context) error {
	p := c.QueryParamOr("page", "")
	l := c.QueryParamOr("limit", "")

	pagination := utils.ResolvePagination(p, l)

	result, err := ah.as.GetAuthors(c.Request().Context(), pagination)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, result)
}

func (ah *AuthorHandler) GetAuthorByID(c *echo.Context) error {
	id, err := utils.ParseUUID(c.Param("id"))
	if err != nil {
		return err
	}

	author, err := ah.as.GetAuthorByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, author)
}

func (ah *AuthorHandler) CreateAuthor(c *echo.Context) error {
	var body types.CreateAuthorRequest
	if err := c.Bind(&body); err != nil {
		return err
	}

	id, err := ah.as.CreateAuthor(c.Request().Context(), body)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusCreated, types.CreateAuthorResult{
		ID: id,
	})
}

func (ah *AuthorHandler) UpdateAuthor(c *echo.Context) error {
	id, err := utils.ParseUUID(c.Param("id"))
	if err != nil {
		return err
	}

	var body types.UpdateAuthorRequest
	if err := c.Bind(&body); err != nil {
		return err
	}

	author, err := ah.as.UpdateAuthor(c.Request().Context(), id, body)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, author)
}

func (ah *AuthorHandler) DeleteAuthor(c *echo.Context) error {
	id, err := utils.ParseUUID(c.Param("id"))
	if err != nil {
		return err
	}

	if err := ah.as.DeleteAuthor(c.Request().Context(), id); err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, map[string]any{
		"status": "success",
	})
}

func (ah *AuthorHandler) GetAuthorBooks(c *echo.Context) error {
	id, err := utils.ParseUUID(c.Param("id"))
	if err != nil {
		return err
	}

	p := c.QueryParamOr("page", "")
	l := c.QueryParamOr("limit", "")

	pagination := utils.ResolvePagination(p, l)

	result, err := ah.as.GetAuthorBooks(c.Request().Context(), id, pagination)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, http.StatusOK, result)
}
