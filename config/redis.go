package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func Redis() *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "starter-redis:6379",
			Password: "",
			DB:       1,
		},
	)
	ping := client.Ping(context.Background())
	if ping.Err() != nil {
		log.Fatalln(ping.Err())
	}
	return client
}
