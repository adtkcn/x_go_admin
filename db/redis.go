package db

import (
	"fmt"
	"x-gin-admin/config"

	"github.com/redis/go-redis/v9"
)

// var ctx = context.Background()
var Redis *redis.Client

func init() {
	var Addr = config.Options.Redis.Addr
	var Password = config.Options.Redis.Password
	var DB = config.Options.Redis.DB

	if Addr == "" {
		fmt.Println("未配置redis.addr")
		return
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password, // no password set
		DB:       DB,       // use default DB
	})

	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// Output: key value
	// key2 does not exist
}
