package main

import (
	"github.com/labstack/echo"
	"github.com/neulhan/gistGoServer/views"
)

func main() {
	e := echo.New()
	e.GET("", views.SchoolMeal)
	e.POST("", views.SchoolMeal)
	e.Logger.Fatal(e.Start(":9090"))
	// views.JSONTest()
}
