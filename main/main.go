package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"../hello"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェアはここ
	// ロガーとか
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング。
	e.GET("/hello", hello.MainPage())

	// サーバー起動
	e.Start(":1323")
}