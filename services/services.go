package services

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/neulhan/gistGoServer/utils"
)

// RequestToWebtoon 웹툰으로 보내는 요청
func RequestToWebtoon() (results []map[string]string) {
	url := "htts://comic.naver.com/webtoon/weekdayList.nhn?week="

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
