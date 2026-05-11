package handlers

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v5"

	"github.com/kylerequez/go-echo-library/src/repositories"
	"github.com/kylerequez/go-echo-library/src/services"
)

func Init(e *echo.Echo, db *sql.DB) error {
	queries := repositories.New(db)
	as := services.NewAuthorService(queries)
	ah := NewAuthorHandler(as)

	bs := services.NewBookService(queries)
	bh := NewBookHandler(bs)

	v1 := "/api/v1"
	authorRoutes := e.Group(fmt.Sprintf("%s/authors", v1))
	authorRoutes.GET("", ah.GetAuthors)
	authorRoutes.POST("", ah.CreateAuthor)
	authorRoutes.GET("/:id", ah.GetAuthorByID)
	authorRoutes.PATCH("/:id", ah.UpdateAuthor)
	authorRoutes.DELETE("/:id", ah.DeleteAuthor)
	authorRoutes.GET("/:id/books", ah.GetAuthorBooks)

	bookRoutes := e.Group(fmt.Sprintf("%s/books", v1))
	bookRoutes.GET("", bh.GetAllBooks)
	bookRoutes.GET("/:id", bh.GetBookByID)
	bookRoutes.POST("", bh.CreateBook)
	bookRoutes.PATCH("/:id", bh.UpdateBook)
	bookRoutes.DELETE("/:id", bh.DeleteBook)

	return nil
}
