package dbdata

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis(host, port string) (*redis.Client, error) {
	// parameter is static, no need to make it parameterizable
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	cmd := client.Ping(ctx)
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	return client, nil
}
