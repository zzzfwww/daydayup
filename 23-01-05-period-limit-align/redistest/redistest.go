package redistest

import (
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)


type PlaceholderType = struct{}

// CreateRedis returns an in process redis.Redis.
func CreateRedis() (r *redis.Client, clean func(), err error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}

	return redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	}), func() {
		ch := make(chan PlaceholderType)

		go func() {
			mr.Close()
			close(ch)
		}()

		select {
		case <-ch:
		case <-time.After(time.Second):
		}
	}, nil
}
