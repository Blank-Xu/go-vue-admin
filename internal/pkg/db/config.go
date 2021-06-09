package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type Config struct {
	EngineName             string       `yaml:"EngineName"`
	DriverName             string       `yaml:"DriverName"`
	DBName                 string       `yaml:"DBName"`
	Host                   string       `yaml:"Host"`
	Port                   string       `yaml:"Port"`
	Username               string       `yaml:"Username"`
	Password               string       `yaml:"Password"`
	Charset                string       `yaml:"Charset"`
	LogLevel               log.LogLevel `yaml:"LogLevel"`
	ConnMaxLifetimeMinutes int          `yaml:"ConnMaxLifetimeMinutes"`
	MaxIdleConns           int          `yaml:"MaxIdleConns"`
	MaxOpenConns           int          `yaml:"MaxOpenConns"`
	ShowSql                bool         `yaml:"ShowSql"`
	Connect                bool         `yaml:"Connect"`
}

func (p *Config) NewEngine(logger *zerolog.Logger) (*xorm.Engine, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		p.Username, p.Password, p.Host, p.Port, p.DBName, p.Charset)

	engine, err := xorm.NewEngine(p.DriverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("database[%s] engine create failed: %v", p.DBName, err)
	}

	engine.SetLogger(NewLogger(logger, p.DBName, p.LogLevel))

	if p.Connect {
		if err = engine.Ping(); err != nil {
			return nil, fmt.Errorf("database[%s] connect failed: %v", p.DBName, err)
		}
		logger.Info().Msgf("database[%s] connect success", p.DBName)
	}

	engine.SetConnMaxLifetime(time.Minute * time.Duration(p.ConnMaxLifetimeMinutes))
	engine.SetMaxIdleConns(p.MaxIdleConns)
	engine.SetMaxOpenConns(p.MaxOpenConns)
	engine.ShowSQL(p.ShowSql)

	return engine, nil
}
