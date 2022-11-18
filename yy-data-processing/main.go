package main

import (
	"log"
	"yy-data-processing/common/config"
	"yy-data-processing/common/database"
	"yy-data-processing/router"
)

func main() {
	// 初始化路由
	routers := router.InitRouters()

	// 初始化mysql
	db := database.InitMysql()
	if db != nil {
		log.Println("db连接成功")
	}
	defer db.Close()

	port := config.Config.Port
	if err := routers.Run(":" + port); err != nil {
		log.Fatalf("启动服务失败: ", err)
	}
}
