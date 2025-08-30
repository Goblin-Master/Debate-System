package model

import (
	"Debate-System/pkg/gormx"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		User{},
	)
	fmt.Println("数据库迁移成功")
	return err
}
func InitDB(cfg gormx.Config, logger logger.Interface) *gorm.DB {
	db, err := gormx.Open(cfg, logger)
	if err != nil {
		panic(err)
	}
	err = autoMigrate(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库初始化成功")
	return db
}
