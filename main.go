package main

import (
    "./handler"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())

    e.GET("/md", handler.Default)
    e.POST("/md", handler.Default)

    e.POST("api/md/file", handler.File)

    e.Start(":8080")
}
