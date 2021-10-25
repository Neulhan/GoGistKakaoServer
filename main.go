package main

import (
	"github.com/Neulhan/GoGistKakaoServer/views"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/:id", views.SchoolMeal)

	e.Logger.Fatal(e.Start(":9090"))
}
