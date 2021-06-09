package dao

import (
	"fmt"

	"xorm.io/xorm"
)

var (
	dbRead  = &xorm.Engine{}
	dbWrite = &xorm.Engine{}
)

func SetEngine(readEngine, writeEngine *xorm.Engine) error {
	if readEngine == nil || writeEngine == nil {
		return fmt.Errorf("engine is nil")
	}
	dbRead = readEngine
	dbWrite = writeEngine
	return nil
}

type dao struct{}

func (*dao) DatabaseReadEngine() *xorm.Engine {
	return dbRead
}

func (*dao) DatabaseWriteEngine() *xorm.Engine {
	return dbWrite
}
