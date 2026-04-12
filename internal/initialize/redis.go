package initialize

import (
	"context"
	"fmt"
	"go-ecomerce-backend-api/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis: %v", zap.Error(err))
	}

	fmt.Println("Connected to Redis")
	global.Rdb = rdb
	RedisExample()
}

func RedisExample() {
	global.Rdb.Set(ctx, "key", "value", 0)
	val, err := global.Rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("Error getting value:", err)
	} else {
		fmt.Println("Value:", val)
	}

	value, err := global.Rdb.Get(ctx, "key2").Result()
	if err != nil {
		fmt.Println("Error getting value:", zap.Error(err))
		return
	}

	global.Logger.Info("Value from Redis", zap.String("key2", value))
}
