package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/kylerequez/go-echo-library/src/database"
	"github.com/kylerequez/go-echo-library/src/handlers"
	"github.com/kylerequez/go-echo-library/src/utils"
)

func Run() error {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	if err := utils.LoadEnv(".env.local"); err != nil {
		return err
	}

	if err := database.Connect(); err != nil {
		return err
	}
	defer database.Disconnect()

	e.GET("/ping", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "pong",
		})
	})

	if err := handlers.Init(e, database.DB); err != nil {
		return err
	}

	port, err := utils.GetEnv("SERVER_PORT")
	if err != nil {
		return err
	}
	return e.Start(fmt.Sprintf(":%s", port))
}
