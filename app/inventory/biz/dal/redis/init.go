package redis

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/conf"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	RS     *redsync.Redsync
)

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := Client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	//初始化分布式锁
	pool := goredis.NewPool(Client)
	RS = redsync.New(pool)
}
