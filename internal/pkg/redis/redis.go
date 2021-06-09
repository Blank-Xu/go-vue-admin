package redis

import (
	"github.com/go-redis/redis"
)

const separator = ":"

var (
	Nil = redis.Nil

	DefaultClient = &redis.Client{}
)
