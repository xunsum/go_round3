package routers

import (
	"github.com/gin-gonic/gin"
	"go_test_project/controllers/contentControllers"
	"go_test_project/middlewares"
)

func ContentRoutersInit(r *gin.Engine) {
	contentRouter := r.Group("/content")
	contentRouter.Use(middlewares.AuthJWT())

	//新增待办
	contentRouter.POST("/addItem", contentControllers.AddItem)

	//设置待办状态
	contentRouter.PUT("/setItemState", contentControllers.SetItemState)

	//设置所有待办状态
	contentRouter.PUT("/setAllItemState", contentControllers.SetAllItemState)

	//过滤 批量 获取待办
	contentRouter.GET("/getItems", contentControllers.GetItems)

	//搜索待办
	contentRouter.GET("/searchItems", contentControllers.SearchItems)

	//删除待办
	contentRouter.DELETE("/deleteItem", contentControllers.DeleteItem)

	//批量删除待办
	contentRouter.DELETE("/deleteItems", contentControllers.DeleteItems)

}
