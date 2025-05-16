package gormx

import (
	"Debate-System/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 使用说明书
// 在config.go的结构体中定义一个xgorm.Mysql或者是xgorm.Postgres，然后yaml文件中写对应的配置，最后用MustOpen方法打开数据库，就能用了

type Config interface {
	getDSN() string
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (cfg Mysql) getDSN() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
}

func (cfg Postgres) getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.Port)
}

// Gorm链接，默认是静默模式
func Open(cfg Config, l logger.Interface) (*gorm.DB, error) {
	dsn := cfg.getDSN()
	var open gorm.Dialector
	switch cfg.(type) {
	case Mysql:
		open = mysql.Open(dsn)
	case Postgres:
		open = postgres.Open(dsn)
	}
	if l == nil {
		l = logger.Default.LogMode(logger.Silent)
	}
	db, err := gorm.Open(open, &gorm.Config{Logger: l})
	return db, err
}
func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.User{},
	)
	fmt.Println("数据库迁移成功")
	return err
}
func InitDB(cfg Config, logger logger.Interface) *gorm.DB {
	db, err := Open(cfg, logger)
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
