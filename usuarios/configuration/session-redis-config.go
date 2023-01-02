package configuration

import (
	"github.com/go-redis/redis/v8"
)

var InstanceRedis *redis.Client

func IniciarRedis() {

	InstanceRedis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})

}
