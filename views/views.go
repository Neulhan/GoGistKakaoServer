package views

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/neulhan/gistGoServer/services"
	"github.com/neulhan/gistGoServer/template"
)

// SchoolMeal 학생식당  크롤링
func SchoolMeal(c echo.Context) error {

	resultWebtoon := services.RequestToWebtoon()
	kakaoResponse := template.JSONResMaker(resultWebtoon)
	return c.JSON(http.StatusOK, kakaoResponse)
}
