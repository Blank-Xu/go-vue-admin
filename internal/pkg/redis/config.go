package redis

import (
	"crypto/tls"

	"github.com/go-redis/redis"
)

type Config struct {
	Addr     string
	Password string
	DB       int
	UseTLS   bool
}

func (p *Config) Init() error {
	client, err := p.NewClient()
	if err != nil {
		return err
	}
	DefaultClient = client
	return nil
}

func (p *Config) NewClient() (*redis.Client, error) {
	options := redis.Options{
		Addr:     p.Addr,
		Password: p.Password,
		DB:       p.DB,
	}
	if p.UseTLS {
		options.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client := redis.NewClient(&options)
	_, err := client.Ping().Result()

	return client, err
}
