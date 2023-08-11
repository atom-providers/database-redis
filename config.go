package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "Redis"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Config struct {
	ClusterName string
	Addresses   []string
	Password    string
	DB          uint
}

func (c *Config) IsClusterMode() bool {
	return c.ClusterName == ""
}

func (c *Config) ToRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.Addresses[0],
		Password: c.Password,
		DB:       int(c.DB),
	}
}

func (c *Config) ToRedisClusterOptions() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		ClientName: c.ClusterName,
		Addrs:      c.Addresses,
		Password:   c.Password,
	}
}
