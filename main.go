package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"./services"
	"./template"
	"./utils"
	"./views"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

func main() {
	redisClient := utils.LoadRedis()
	e := echo.New()
	e.GET("/w", views.Webtoon)
	e.POST("/w", views.Webtoon)

	e.GET("/:id", func(c echo.Context) error {

		var resultGIST []map[string]string
		val, err := redisClient.Get("resultGIST").Result()

		if err == redis.Nil {
			// key 에 대응되는 값이 존재하지 않을 경우
			resultGIST = services.RequestToGIST()
			jsonResultGIST, _ := json.Marshal(resultGIST)

			err2 := redisClient.Set("resultGIST", jsonResultGIST, time.Hour).Err()
			fmt.Println(err2)

		} else if err != nil {
			// 에러 발생
			fmt.Println(err)
		} else {
			// key 에 대응되는 값을 찾았을 경우
			json.Unmarshal([]byte(val), &resultGIST)
		}

		id, _ := strconv.Atoi(c.Param("id"))
		kakaoResponse := template.JSONResMaker(resultGIST[id : id+1])
		return c.JSON(http.StatusOK, kakaoResponse)
	})
	e.POST("/:id", views.SchoolMeal)

	e.Logger.Fatal(e.Start(":9090"))

	// utils.SlackSender(config.ErrorWebhook, "g")
	// views.JSONTest()
}
