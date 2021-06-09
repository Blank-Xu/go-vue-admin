package dao

import (
	"go-vue-admin/internal/models"
	"go-vue-admin/internal/pkg/db"
)

type SysUser struct {
	dao
	db.Model
	models.SysUser
}

func (*SysUser) TableName() string {
	return "sys_admin"
}
