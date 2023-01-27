package utils

import (
	"github.com/go-redis/redis/v8"
)

var Redis = redis.NewClient(&redis.Options{
	Addr:     Config.RedisAddress,
	Password: "", // no password set
	DB:       0,  // use default DB
})
