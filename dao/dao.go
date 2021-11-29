package dao

import (
	"aurora/core"
	"gorm.io/gorm"
)

type Dao struct {
	DB *gorm.DB
}

func New() *Dao {
	d := &Dao{
		DB:core.InitMysql(),
	}

	return d
}
