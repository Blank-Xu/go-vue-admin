package http

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Config struct {
	IP                  string `yaml:"IP"`
	Port                int    `yaml:"Port"`
	ReadTimeoutSeconds  int    `yaml:"ReadTimeoutSeconds"`
	WriteTimeoutSeconds int    `yaml:"WriteTimeoutSeconds"`
	IdleTimeoutSeconds  int    `yaml:"IdleTimeoutSeconds"`
	MaxHeaderMB         int    `json:"MaxHeaderMB"`
}

func (p *Config) Addr() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func (p *Config) NewHttpServer(router *http.ServeMux, logger *log.Logger) *http.Server {
	return &http.Server{
		Addr:           p.Addr(),
		Handler:        router,
		ReadTimeout:    time.Second * time.Duration(p.ReadTimeoutSeconds),
		WriteTimeout:   time.Second * time.Duration(p.WriteTimeoutSeconds),
		IdleTimeout:    time.Second * time.Duration(p.IdleTimeoutSeconds),
		MaxHeaderBytes: p.MaxHeaderMB * 1024 * 1024,
		ErrorLog:       logger,
	}
}
