package hello

import (
	"github.com/labstack/echo"
	"net/http"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		//cがコンテキストで、Request, Responseを色々できる
		return c.String(http.StatusOK, "Hello World")
	}
}