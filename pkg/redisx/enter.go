package redisx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host     string
	Port     int
	DB       int
	Password string
	Enable   bool
}

func (r *Redis) DSN() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

func InitRedis(config Redis) redis.Cmdable {
	if !config.Enable {
		return nil
	}
	client := redis.NewClient(&redis.Options{
		Network:         "",
		Addr:            config.DSN(),
		Dialer:          nil,
		OnConnect:       nil,
		Username:        "",
		Password:        config.Password,
		DB:              config.DB,
		MaxRetries:      0,
		MinRetryBackoff: 0,
		MaxRetryBackoff: 0,
		DialTimeout:     0,
		ReadTimeout:     0,
		WriteTimeout:    0,
		PoolFIFO:        false,
		PoolSize:        1000,
		MinIdleConns:    1,
		PoolTimeout:     0,
		TLSConfig:       nil,
		Limiter:         nil,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(fmt.Sprintf("redis init error:%v", err))
	}
	return client
}
