package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client

func ZAddDemo() {
	// key
	zsetKey := "language_rank"
	// value
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// ZADD
	rdb.ZAdd(ctx, zsetKey, languages...)
	fmt.Println("zadd success")
	// 把Golang的分数加10
	newScore, _ := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, _ = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行命令获取结果
	val, err := rdb.Get(ctx, "lsy").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("key not exit")
		}
	}
	fmt.Println(val, err)
	ZAddDemo()
}
