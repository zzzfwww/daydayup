package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
}

func NewRedisClint() *Client {
	return &Client{Client: redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
}

func (c *Client) ExampleTest(ctx context.Context) {
	err := c.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := c.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := c.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func (c *Client) SetRedis(ctx context.Context, key, value string) {
	err := c.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (c *Client) GetRedis(ctx context.Context, key string) string {
	val, err := c.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
