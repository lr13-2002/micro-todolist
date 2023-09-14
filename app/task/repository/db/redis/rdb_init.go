package redis

import (
	"context"
	"fmt"
	"log"
	"micro-todolist/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRDB() {
	host := config.RDbHost
	port := config.RDbPort
	dsn := fmt.Sprintf("%s:%s", host, port)
	rdb, err := Redis(dsn)

	if err != nil {
		panic(err)
	}
	redisClient = rdb
}

func Redis(connString string) (*redis.Client, error) {
	log.Println(connString)
	rdb := redis.NewClient(&redis.Options{
		Addr:     connString,
		Password: "",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e := rdb.Set(ctx, "12", "12", 100*time.Second).Err()

	if e != nil {
		panic(e)
	}

	result, er := rdb.Get(ctx, "12").Result()
	if er != nil {
		panic(er)
	}
	log.Println(result)
	_, err := rdb.Ping(ctx).Result()
	return rdb, err
}
