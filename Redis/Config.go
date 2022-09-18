package Redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var Redis *redis.Client
var ctx context.Context

func init() {

	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx = context.Background()

	if _, err := Redis.Ping(ctx).Result();err != nil {
		fmt.Println("------------------- server redis doesn't exists -------------------")
	}
}
