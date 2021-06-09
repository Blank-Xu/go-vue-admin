package cmd

import (
	"xorm.io/xorm"

	"go-vue-admin/internal/app/dao"
	"go-vue-admin/internal/pkg/log"
)

func loadDatabaseEngine() error {
	var (
		readEngine, writeEngine *xorm.Engine
		err                     error
	)
	for _, cfg := range Config.Database {
		switch cfg.EngineName {
		case "ReadEngine":
			readEngine, err = cfg.NewEngine(&log.Logger)
			if err != nil {
				log.Error().Msgf("load database read engine failed, err: %v", err)
				return err
			}
			log.Info().Msg("load database read engine success")
		case "WriteEngine":
			writeEngine, err = cfg.NewEngine(&log.Logger)
			if err != nil {
				log.Error().Msgf("load database write engine failed, err: %v", err)
				return err
			}
			log.Info().Msg("load database write engine success")
		}
	}

	return dao.SetEngine(readEngine, writeEngine)
}
