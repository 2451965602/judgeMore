package cache

import (
	"github.com/redis/go-redis/v9"
	client "judgeMore/pkg/ba"
	"judgeMore/pkg/constants"
)

var userCa *redis.Client
var eventCa *redis.Client

func Init() {
	var err error
	userCa, err = client.NewRedisClient(constants.RedisDBUser)
	if err != nil {
		panic(err)
	}
	eventCa, err = client.NewRedisClient(constants.RedisDBEvent)
	if err != nil {
		panic(err)
	}
}
