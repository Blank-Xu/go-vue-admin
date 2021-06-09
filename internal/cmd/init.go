package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"go-vue-admin/configs"
)

var (
	// Config .
	Config = &config{}
)

// Init  .
func Init(ctx context.Context) {
	// read config
	err := yaml.Unmarshal(configs.AppFile, Config)
	if err != nil {
		panic(fmt.Errorf("decode default config file failed, err: %v", err))
	}

	if confFile != "" {
		bs, err := os.ReadFile(confFile)
		if err != nil {
			panic(fmt.Errorf("read config file[%s] failed, err: %v", confFile, err))
		}
		if err = yaml.Unmarshal(bs, Config); err != nil {
			panic(fmt.Errorf("decode config file[%s] failed, err: %v", confFile, err))
		}
		log.Println("read config success")
	}

	// load config
	if err = Config.Log.Init(ctx); err != nil {
		panic(fmt.Errorf("load log failed, err: %v", err))
	}
	if err = loadDatabaseEngine(); err != nil {
		panic(fmt.Errorf("load database engine failed, err: %v", err))
	}

}
