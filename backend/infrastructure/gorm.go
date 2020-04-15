package infrastructure

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.FetchEnvValue("MYSQL_USER", "BambooTuna"),
		config.FetchEnvValue("MYSQL_PASS", "pass"),
		config.FetchEnvValue("MYSQL_HOST", "127.0.0.1"),
		config.FetchEnvValue("MYSQL_PORT", "3306"),
		config.FetchEnvValue("MYSQL_DATABASE", "letustalk"),
	)
	db, err := gorm.Open("mysql", connect)

	if err != nil {
		panic(err.Error())
	}
	return db
}
