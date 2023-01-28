package utils

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var Redis = redis.NewClient(&redis.Options{
	Addr:     Config.RedisAddress,
	Password: "", // no password set
	DB:       0,  // use default DB
	OnConnect: func(ctx context.Context, cn *redis.Conn) error {
		log.Debug("Connected to redis")

		return nil
	},
})

var Cache = cache.New(&cache.Options{
	Redis: Redis,
})
