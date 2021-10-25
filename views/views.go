package views

import (
	"net/http"
	"strconv"

	"github.com/Neulhan/GoGistKakaoServer/services"
	"github.com/Neulhan/GoGistKakaoServer/template"
	"github.com/labstack/echo"
)

// SchoolMeal 학생식당  크롤링
func SchoolMeal(c echo.Context) error {
	resultGIST := services.RequestToGIST()
	id, _ := strconv.Atoi(c.Param("id"))
	kakaoResponse := template.JSONResMaker(resultGIST[id : id+1])
	return c.JSON(http.StatusOK, kakaoResponse)
}

// Webtoon 웹툰으로 테스트
func Webtoon(c echo.Context) error {
	resultWebtoon := services.RequestToWebtoon()
	kakaoResponse := template.JSONResMaker(resultWebtoon)
	return c.JSON(http.StatusOK, kakaoResponse)
}
