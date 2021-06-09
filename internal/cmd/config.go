package cmd

import (
	"go-vue-admin/internal/pkg/db"
	"go-vue-admin/internal/pkg/http"
	"go-vue-admin/internal/pkg/log"
	"go-vue-admin/internal/pkg/redis"
)

// VERSION  .
const VERSION = "0.1.0"

type config struct {
	Server   *http.Config    `json:"Server"`
	Log      *log.Config     `yaml:"Log"`
	Redis    []*redis.Config `yaml:"Redis"`
	Database []*db.Config    `yaml:"Database"`
}
