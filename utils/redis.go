package utils

import "github.com/go-redis/redis"

// LoadRedis redis 서버를 로드하는 함수
func LoadRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis 서버 주소 (redis 의 디폴트 포트 6379 로컬호스트)
		Password: "",               // redis 비밀번호
		DB:       0,                // redis DB 번호 선택
	})
	return redisClient
}
