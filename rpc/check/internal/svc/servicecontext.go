package svc

import (
	"bookstore/rpc/check/internal/config"
	"bookstore/rpc/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := "root:root123@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.BookStore{})

	return &ServiceContext{
		Config: c,
		Db: db,
	}
}
