package services

import (
	"encoding/json"
	"github.com/Neulhan/GoGistKakaoServer/template"
	"net/http"

	"github.com/Neulhan/GoGistKakaoServer/config"
	"github.com/Neulhan/GoGistKakaoServer/utils"
	"github.com/PuerkitoBio/goquery"
)

type jsonImageSrc struct {
	ID       string
	Src      string
	LinkUser string
	Link     string
}

// RequestToGIST 지스트 학식 스크래핑하는 함수입니다.
func RequestToGIST() (cards []template.KakaoCard) {
	res, err := http.Get(config.GISTBlogURL)
	utils.CheckError(err)
	utils.CheckResponse(res)
	doc, qErr := goquery.NewDocumentFromReader(res.Body)
	utils.CheckError(qErr)
	defer res.Body.Close()

	doc.Find(".se-module-image-link").Each(func(idx int, s *goquery.Selection) {
		jsonSrc, _ := s.Attr("data-linkdata")
		jsonBtye := []byte(jsonSrc)
		var j jsonImageSrc
		json.Unmarshal(jsonBtye, &j)
		card := template.SimpleImageCard{Text: j.Src  + "?type=w800"}
		cards = append(cards, card)
	})
	return
}
