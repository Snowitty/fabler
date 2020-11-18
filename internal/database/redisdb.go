package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/snowitty/fabler/conf"
	"log"
)

var RDB *redis.Client
var ctx = context.Background()

func init(){
	RDB = redis.NewClient(&redis.Options{
		Addr: conf.Config().Redis.Addr,
		Password: conf.Config().Redis.Password,
		DB: conf.Config().Redis.Db,
	})

	if dbsize, err := RDB.DBSize(ctx).Result(); err != nil{
		log.Panicln("[redis]: error", err)
		panic("failed to connect redis")
	}else {
		log.Println("[redis]: dbsize", dbsize)
	}
}