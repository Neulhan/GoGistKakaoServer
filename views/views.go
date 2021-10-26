package views

import (
	"net/http"
	"strconv"

	"github.com/Neulhan/GoGistKakaoServer/services"
	"github.com/Neulhan/GoGistKakaoServer/template"
	"github.com/labstack/echo"
)


var extraText = "식단표 이미지가 보이지 않으시거나 내용이 다를 경우 링크를 눌러 확인해주세요 :)"
var nameList = []string{"-1학1층 식당", "-1학2층 식당", "-2학1층 식당"}
var linkList = []string{"https://www.gist.ac.kr/kr/html/sub05/050601.html", "https://www.gist.ac.kr/kr/html/sub05/050603.html", "https://www.gist.ac.kr/kr/html/sub05/050602.html"}

// SchoolMeal 학생식당  크롤링
func SchoolMeal(c echo.Context) error {
	resultGIST := services.RequestToGIST()
	id, _ := strconv.Atoi(c.Param("id"))
	card2 := template.SimpleTextCard{Text: extraText + "\n" + nameList[id - 1] + "\n" + linkList[id - 1]}
	cardList := append(resultGIST[id : id + 1], card2)
	kakaoResponse := template.JSONResMaker(cardList)
	return c.JSON(http.StatusOK, kakaoResponse)
}
