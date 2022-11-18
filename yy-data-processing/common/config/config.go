package config

import (
	"gopkg.in/ini.v1"
	"log"
	"yy-data-processing/model"
)

var Cfg *ini.File
var Config model.Config

func init() {
	var err error
	Cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config.ini': %v", err)
	}

	loadApp()
	loadMysql()
}

func loadApp() {
	Config.Port = Cfg.Section("app").Key("Port").String()
}

func loadMysql() {
	Config.MysqlUrl = Cfg.Section("mysql").Key("Url").String()
}
