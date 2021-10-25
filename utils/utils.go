package utils

import (
	"bytes"
	"fmt"
	"net/http"

	"strconv"
	"strings"

	"../config"
	"github.com/PuerkitoBio/goquery"
)

// CheckError 에러 체크
func CheckError(e error) {
	if e != nil {
		err := errorToSlackMessage(e)
		// SlackSender(config.ErrorWebhook, err)
		fmt.Println(err)
	}
}

// CheckResponse 에러 체크
func CheckResponse(res *http.Response) {
	if res.StatusCode >= 400 {
		err := responseErrorToSlackMessage(res)
		SlackSender(config.ErrorWebhook, err)
	}
}

// SlackSender 슬랙으로 알림을 보내는 함수
func SlackSender(u string, s string) {

	reqBody := bytes.NewBufferString(s)

	res, err := http.Post(u, "text/plain", reqBody)

	fmt.Println(res.StatusCode)
	fmt.Println(err)
}

func errorToSlackMessage(e error) string {
	// errMsg := `echo\: http\: panic serving 127.0.0.1\:34176\: runtime error\: invalid memory address or nil pointer dereference
	// goroutine 5 [running]\:`
	errMsg := e.Error()
	// errMsg := "http://127.0.0.1:10108/"
	r1 := strings.NewReplacer(`"`, `\"`)
	msg := `
	{
		"attachments":[
		   {
			  "fallback":"gist kakao 채널에 에러 발생!",
			  "pretext":"gist kakao 채널에 에러 발생!",
			  "color":"#FF0000",
			  "fields":[
				 {
					"title":"사령관님 문제상황이 발생했습니다!!",
					"value": "` + r1.Replace(errMsg) + `",
					"short":false
				 }
			  ]
		   }
		]
	 }`

	r2 := strings.NewReplacer(`&`, `&amp`, `<`, `&lt`, `>`, `&gt`)
	result := r2.Replace(msg)
	fmt.Println(result)
	return result
}

func responseErrorToSlackMessage(e *http.Response) string {
	doc, _ := goquery.NewDocumentFromReader(e.Body)
	errMsg := "[" + strconv.Itoa(e.StatusCode) + " error]: " + doc.Text()

	defer e.Body.Close()
	r1 := strings.NewReplacer(`"`, `\"`)
	msg := `
	{
		"attachments":[
		   {
			  "fallback":"gist kakao 채널에 에러 발생!",
			  "pretext":"gist kakao 채널에 에러 발생!",
			  "color":"#FFA500",
			  "fields":[
				 {
					"title":"사령관님 문제상황이 발생했습니다!!",
					"value": "` + r1.Replace(errMsg) + `",
					"short":false
				 }
			  ]
		   }
		]
	 }`

	r2 := strings.NewReplacer(`&`, `&amp`, `<`, `&lt`, `>`, `&gt`)
	result := r2.Replace(msg)
	fmt.Println(result)
	return result
}
