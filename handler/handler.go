package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {     //cがコンテキストで、Request, Responseを色々する
		return c.String(http.StatusOK, "Hello World")
	}
}