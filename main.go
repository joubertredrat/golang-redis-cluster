package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("0")

	rdb := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs: []string{"127.0.0.1:6379", "127.0.0.1:6380", "127.0.0.1:6381"},
		},
	)

	fmt.Println("1")

	rdb.Ping(context.Background())

	fmt.Println("2")

	client := cache.New(
		&cache.Options{
			Redis:      rdb,
			LocalCache: cache.NewTinyLFU(1000, time.Second),
		},
	)

	fmt.Println("3")

	for i := 0; i < 3; i++ {
		err := client.Set(
			&cache.Item{
				Key:   fmt.Sprintf("key_%d", i),
				Value: fmt.Sprintf("key_%d", i),
			},
		)
		fmt.Println(err)
	}

	fmt.Println("4")
	fmt.Println("end")
}
