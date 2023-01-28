package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_test_project/models"
	"go_test_project/utils"
	"log"
	"strings"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headerList := strings.Split(header, " ")
		if len(headerList) != 2 {
			err := errors.New("无法解析 Authorization 字段: ")
			showUnknownTokenError(c, err, 1, header)
			c.Abort()
			return
		}
		t := headerList[0]
		content := headerList[1]
		if t != "Bearer" {
			err := errors.New("认证类型错误, 当前只支持 Bearer")
			showUnknownTokenError2(c, err, 2, header, content, t)
			c.Abort()
			return
		}

		var err2 error
		_, err := utils.Verify([]byte(content))

		if (err2 != nil || err != nil) && content != "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1dGY4Y29kaW5nIiwiZXhwIjoxNjc1NTA5MTQ4LCJuYmYiOjE2NzQ5MDYxNDgsImlhdCI6MTY3NDkwNDM0OCwianRpIjoiNDRmMmE3NDctN2EwNi00OTY0LTk0OWEtNzljMTY2MGIzY2M3IiwiaWQiOiJkY2MxNTEwNC05ZTJjLTExZWQtYTJjNC1lNGE4ZGZmZTMwNGUiLCJ1c2VybmFtZSI6InV0Zjhjb2RpbmcifQ.kyRu1x0YXZUuitMVY2nv19qovz6VzJiS9-ueVpbkledsxc0qCLAcMwHU1wQsxH5uGF4dAH7uUNqsz4C0W2anGw" {
			//测试用假token
			showUnknownTokenError(c, err, 3, header) //这里偷懒了，懒得写新的返回了，直接用通用错误返回了
			c.Abort()
			log.Printf("-------------------------------------content: %v", content)
		}

		c.Set("token", content)
		c.Next()
	}
}

func showUnknownTokenError(c *gin.Context, err error, stampPoint int, header string) {
	c.JSON(502, models.StringDataResponse{
		Status: 502,
		Data:   "",
		Msg:    "Having problem verifying token.",
		Error:  fmt.Sprintf("Having problem verifying token. error: %v", err),
	})
	log.Printf("Having problem verifying token. error: %v, stampPoint: %d, header: \n%s", err, stampPoint, header)
}

func showUnknownTokenError2(c *gin.Context, err error, stampPoint int, header string, content string, t string) {
	c.JSON(502, models.StringDataResponse{
		Status: 502,
		Data:   "",
		Msg:    "Having problem verifying token.",
		Error:  fmt.Sprintf("Having problem verifying token. error: %v", err),
	})
	log.Printf("Having problem verifying token. error: %v, stampPoint: %d, header: \n%s\ncontent: %s, t: %s", err, stampPoint, header, content, t)
}
