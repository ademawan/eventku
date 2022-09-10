package middlewares

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func BusicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}

	return false, errors.New("bukan admin")
}
