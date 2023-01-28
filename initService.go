package main

import (
	"github.com/gin-gonic/gin"
	"go_test_project/routers"
)

func main() {

	r := gin.Default() // 使用Default创建路由

	//api v1 路由
	r.Group("/api/v1")

	//用户路由
	routers.UserRoutersInit(r)

	//内容路由
	routers.ContentRoutersInit(r)

	_ = r.Run(":8000")
}
