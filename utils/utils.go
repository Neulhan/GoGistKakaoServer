package utils

import (
	"log"
	"net/http"
)

// CheckError 에러 체크
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CheckResponse 에러 체크
func CheckResponse(res *http.Response) {
	if res.StatusCode >= 400 {
		log.Fatal(res.StatusCode)
	}
}
