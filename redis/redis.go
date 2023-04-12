package redis

import (
	"github.com/go-redis/redis"
)

type Option struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool-size"`
	Timeout  int    `yaml:"timeout"`
}

func New(opt *Option) (*redis.Client, error) {
	var (
		err error
		rdb *redis.Client
	)

	rdb = redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
		PoolSize: opt.PoolSize,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return nil, err
	}
	return rdb, err
}
