package main

import (
	"web/handler"
	"web/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/static", "static")

	e.Renderer = utils.NewRenderer()

	e.GET("/", handler.Home)

	e.POST("/browse", handler.Set_dataset)
	e.POST("/upload", handler.Browse_dataset)
	e.POST("/graph", handler.Create_plotter)
	e.GET("/data", handler.Data)

	e.Logger.Fatal(e.Start(":8000"))
}
