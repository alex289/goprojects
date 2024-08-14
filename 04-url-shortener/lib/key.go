package lib

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func GenerateKey(rdb *redis.Client, ctx context.Context) *string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars); j++ {
			key := string(chars[i]) + string(chars[j])
			_, err := rdb.Get(ctx, key).Result()

			if err != nil {
				return &key
			}
		}
	}

	return nil
}
