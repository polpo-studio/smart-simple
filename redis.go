package simple

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var (
	client *redis.Client
)

func OpenRedis(option *redis.Options) (err error) {
	c := redis.NewClient(option)
	pong, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Info("redis connect ping response:", pong)
	} else {
		client = c
	}
	return
}

func Redis() *redis.Client {
	return client
}

func CloseRedis() {
	if client == nil {
		return
	}
	if err := client.Close(); nil != err {
		log.Errorf("Disconnect from redis failed: %s", err.Error())
	}
}
