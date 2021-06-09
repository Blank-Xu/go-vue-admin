package db

import (
	"xorm.io/xorm"
)

type Interface interface {
	TableName() string
	DatabaseReadEngine() *xorm.Engine
	DatabaseWriteEngine() *xorm.Engine
}
