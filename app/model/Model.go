package model

import (
	"fmt"
	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
	"sync"
)

type Model struct {
	*gorm.DB
}

var (
	db *gorm.DB
	once sync.Once
)



func InitDB(cfg *ini.File) {
	once.Do(func() {
		driver := strings.ToLower(cfg.Section(ini.DefaultSection).Key("DB_CONNECTION").MustString("mysql"))
		host := cfg.Section(ini.DefaultSection).Key("DB_HOST").MustString("127.0.0.1")
		port := cfg.Section(ini.DefaultSection).Key("DB_PORT").MustString("3306")
		database := cfg.Section(ini.DefaultSection).Key("DB_DATABASE").MustString("go-chats")
		username := cfg.Section(ini.DefaultSection).Key("DB_USERNAME").MustString("root")
		password := cfg.Section(ini.DefaultSection).Key("DB_PASSWORD").MustString("123456")

		switch driver {
		case "mysql":
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
			var err error
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatalln(fmt.Errorf("Database connection failed error:  %v", err))
			}
			db.AutoMigrate(&User{}) //自动迁移
			new(Model).DB = db
		case "sqlserver":

		case "postgres", "postgre", "postgresql":

		default:
			log.Fatalln(fmt.Errorf("%v SQL database type is not supported", driver))
		}
	})
}

func DB() *gorm.DB {
	return db
}