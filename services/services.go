package services

import (
	"encoding/json"
	"net/http"

	"../config"
	"../utils"
	"github.com/PuerkitoBio/goquery"
)

// RequestToWebtoon 웹툰으로 보내는 요청
func RequestToWebtoon() (results []map[string]string) {
	url := "https://comic.naver.com/webtoon/weekdayList.nhn?week="

	res, rqErr := http.Get(url)
	utils.CheckError(rqErr)
	utils.CheckResponse(res)

	doc, gqErr := goquery.NewDocumentFromReader(res.Body)
	utils.CheckError(gqErr)

	doc.Find("ul.img_list .thumb img").Each(func(idx int, s *goquery.Selection) {
		if idx < 3 {
			src, _ := s.Attr("src")
			result := map[string]string{"imageUrl": src, "altText": "image"}
			results = append(results, result)
		}
	})
	defer res.Body.Close()
	return
}

type jsonImageSrc struct {
	ID       string
	Src      string
	LinkUser string
	Link     string
}

// RequestToGIST 지스트 학식 스크래핑하는 함수입니다.
func RequestToGIST() (images []map[string]string) {
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
		image := map[string]string{"imageUrl": j.Src + "?type=w800", "altText": "image"}
		images = append(images, image)
	})
	return
}
