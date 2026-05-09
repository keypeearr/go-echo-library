package utils

import "github.com/labstack/echo/v5"

func SendResponse(c *echo.Context, status int, payload any) error {
	return c.JSON(status, payload)
}
