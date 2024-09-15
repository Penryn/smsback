package main

import (
	"log"
	"smsback/internal/middleware"
	"smsback/internal/pkg/database"
	"smsback/internal/router"
	"smsback/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db := database.MysqlInit()
	// 初始化dao
	service.ServiceInit(db)

	// 初始化gin
	r := gin.Default()
	r.NoMethod(middleware.HandleNotFound)
	r.NoRoute(middleware.HandleNotFound)
	router.Init(r)
	err := r.Run(":8088")
	if err != nil {
		log.Fatal("Failed to start the server:" + err.Error())
	}
}