package setup

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strings"
)

func NewRedisClient() *redis.Client {
	cacheHost := os.Getenv("REDIS_HOST")
	cachePort := os.Getenv("REDIS_PORT")
	args := os.Args

	for _, arg := range args {
		split := strings.Split(arg, "=")
		switch split[0] {
		case "REDIS_HOST":
			cacheHost = split[1]
			break
		case "REDIS_PORT":
			cachePort = split[1]
			break
		}
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cacheHost, cachePort),
		Password: "",
		DB:       0,
	})

	return client
}
