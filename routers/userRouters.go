package routers

import (
	"github.com/gin-gonic/gin"
	"go_test_project/controllers/userControllers"
)

func UserRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/user")
	{
		//注册
		userRouters.POST("/register", userControllers.Register)

		//登录
		userRouters.GET("/login", userControllers.Login)

		//上传文件
		/*r.POST("/uploadFile", func(c *gin.Context) {
			file, _ := c.FormFile("file")
			c.PostForm("")
			dst := "./" + file.Filename
			err := c.SaveUploadedFile(file, dst)
			if err == nil {
				fmt.Printf("IP %s: upload file success", c.ClientIP())
				c.JSON(http.StatusOK, structs.MsgResponseData{
					Status: http.StatusOK,
					Data:   "",
					Msg:    "upload success",
					Error:  "none",
				})
			} else {
				fmt.Printf("IP %s: file upload failed, error: %e", c.ClientIP(), err)
				c.JSON(http.StatusBadGateway, structs.MsgResponseData{
					Status: http.StatusBadGateway,
					Data:   "",
					Msg:    "upload failed - unknown error",
					Error:  fmt.Sprintf("%e", err),
				})
			}

		})*/

		//获取静态图片
		/*r.GET("/getHelloWorldImage", func(c *gin.Context) {
			c.File("E:/Pictures_Acer/ShaftImages/HELLO WORLD! BY KYOANI.png")
		})*/
	}
}
