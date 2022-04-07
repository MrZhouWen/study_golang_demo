package main

import (
	"gin-answer/controller/favorite"
	"gin-answer/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := make(map[string]string)
	config["log_level"] = "debug"
	logger.InitLogger("console", config)

	//收藏模块路由
	favoriteGroup := router.Group("/api/")
	favoriteGroup.GET("/list", favorite.GetList)

	router.Run(":9090")
}
