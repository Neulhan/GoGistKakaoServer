package views

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type kakaoCard struct {
	SimpleImage map[string]string `json:"simpleImage"`
}
type kakaoTemplate struct {
	Outputs []kakaoCard `json:"outputs"`
}

type kakaoRes struct {
	Version  string        `json:"version"`
	Template kakaoTemplate `json:"template"`
}

func jsonResMaker() (res *kakaoRes) {
	var kakaocards []kakaoCard

	card := map[string]string{"imageUrl": "http://k.kakaocdn.net/dn/83BvP/bl20duRC1Q1/lj3JUcmrzC53YIjNDkqbWK/i_6piz1p.jpg", "altText": "보물상자입니다"}
	fmt.Println(card)
	kakaocard := kakaoCard{SimpleImage: card}
	fmt.Println(kakaocard)

	kakaocards = append(kakaocards, kakaocard)
	template := kakaoTemplate{Outputs: kakaocards}
	fmt.Println(template)

	res = &kakaoRes{Version: "2.0", Template: template}
	fmt.Println(res)
	return
}

// SchoolMeal 학생식당  크롤링
func SchoolMeal(c echo.Context) error {

	// url := "https://comic.naver.com/webtoon/weekdayList.nhn?week="
	// var result = []string{}

	// res, rqErr := http.Get(url)
	// utils.CheckError(rqErr)
	// utils.CheckResponse(res)

	// doc, gqErr := goquery.NewDocumentFromReader(res.Body)
	// defer res.Body.Close()
	// utils.CheckError(gqErr)

	// doc.Find("ul.img_list .thumb img").Each(func(idx int, s *goquery.Selection) {
	// 	src, _ := s.Attr("src")
	// 	result = append(result, src)
	// })
	// resultStr := strings.Join(result, "\n")
	kakaoResponse := jsonResMaker()
	return c.JSON(http.StatusOK, kakaoResponse)
}
