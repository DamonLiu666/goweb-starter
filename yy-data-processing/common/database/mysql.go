package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"yy-data-processing/common/config"
)

// 定义db全局变量
var Db *gorm.DB

func InitMysql() *gorm.DB {
	var err error
	url := config.Config.MysqlUrl
	Db, err = gorm.Open("mysql", url)
	if err != nil {
		log.Fatalf("Connecting database failed: " + err.Error())
	}

	// 设置连接池
	Db.DB().SetMaxIdleConns(5)
	// 最大连接数
	Db.DB().SetMaxOpenConns(10)
	// 禁用后缀加s，结构体对应的表名不用带s
	Db.SingularTable(true)

	return Db
}
