package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/neulhan/gistGoServer/config"
)

// CheckError 에러 체크
func CheckError(e error) bool {
	if e != nil {
		err := errorToSlackMessage(e)
		SlackSender(config.ErrorWebhook, err)
		return true
	}
	return false
}

// CheckResponse 에러 체크
func CheckResponse(res *http.Response) {
	if res.StatusCode >= 400 {
		log.Fatal(res.StatusCode)
	}
}

// SlackSender 슬랙으로 알림을 보내는 함수
func SlackSender(u string, s string) {

	reqBody := bytes.NewBufferString(s)

	res, err := http.Post(u, "text/plain", reqBody)

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	fmt.Println(doc.Text())
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Println(err)
}

func errorToSlackMessage(e error) (msg string) {
	// errMsg := "e.Error()"
	errMsg := e.Error()

	msg = `
	{
		"attachments":[
		   {
			  "fallback":"New open task [Urgent]: <http://url_to_task|Test out Slack message attachments>"
			  "text": ` + errMsg + ` 
			  "pretext":"New open task [Urgent]: <http://url_to_task|Test out Slack message attachments>",
			  "color":"#FF0000",
			  "fields":[
				 {
					"title":"사령관님 오류가 발견되었습니다!!",
					"value":"에러입니다!",
					"short":false
				 }
			  ]
		   }
		]
	}`
	fmt.Println(2222)
	return
}
