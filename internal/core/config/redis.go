package config

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

func (r *Redis) Options() *redis.Options {
	return &redis.Options{
		Addr:         r.Addr,
		Password:     r.Password,
		DB:           r.DB,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     20,
		MinIdleConns: 5,
	}
}
