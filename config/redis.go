package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// ConnectRedis init new redis connection
func ConnectRedis() (*redis.Client, error) {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv(fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))),
		DB:   redisDB,
	})
	return rdb, nil
}
