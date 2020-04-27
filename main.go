package main

import (
	"github.com/labstack/echo"
	"github.com/neulhan/gistGoServer/views"
)

func main() {
	e := echo.New()
	e.GET("/w", views.Webtoon)
	e.POST("/w", views.Webtoon)

	e.GET("/:id", views.SchoolMeal)
	e.POST("/:id", views.SchoolMeal)

	e.Logger.Fatal(e.Start(":10108"))

	// utils.SlackSender(config.ErrorWebhook, "g")
	// views.JSONTest()
}
