package repository

import (
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/rpolnx/go-asynq-poc/internal/configs"
)

func NewCacheClient(cfg *configs.AppConfig) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Username: cfg.Redis.User,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
		PoolSize: 10,
	})
}

func NewCacheServer(cfg *configs.AppConfig) *asynq.Server {
	return asynq.NewServer(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Username: cfg.Redis.User,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
		PoolSize: 10,
	},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		})
}
