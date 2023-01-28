package userControllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go_test_project/models"
	"go_test_project/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func Register(c *gin.Context) {
	newUser := models.User{}
	userName := c.PostForm("userName")
	userPswd := c.PostForm("userPassword")
	userMail := c.PostForm("userMail")
	//检查重名
	var searchOutcome []models.User
	result := utils.DB.Where("user_name = ?", userName).First(&searchOutcome)
	if fmt.Sprintf("%v", result.Error) != "record not found" && result.Error != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Server have problems searching databases.",
			Error:  fmt.Sprintf("Databases search unavailable, err: %v", result.Error),
		})
		log.Printf("Databases search unavailable, err: %v", result.Error)
	} else if userName == "" || userPswd == "" || userMail == "" {
		c.JSON(500, models.StringDataResponse{
			Status: 500,
			Data:   "",
			Msg:    "Null user name or password.",
			Error:  "Illegal input",
		})
		log.Printf("Illegal user name or password: %s, %s, %s \n", userName, userPswd, userMail)
	} else if result.Error == nil {
		c.JSON(500, models.StringDataResponse{
			Status: 500,
			Data:   "",
			Msg:    "The user name has been used.",
			Error:  "Used user name",
		})
		log.Printf("Used user name")
	} else {
		//创建id
		userId, err := uuid.NewUUID()
		if err != nil {
			c.JSON(502, models.StringDataResponse{
				Status: 502,
				Data:   "",
				Msg:    "Having problem generating user id.",
				Error:  fmt.Sprintf("Uuid generation failed error: %v", err),
			})
			log.Printf("Uuid generation failed error: %v", err)
		} else {
			newUser.Id = userId.String()
			newUser.UserName = userName
			newUser.Email = userMail
			//获取时间
			newUser.GenerateTime = time.Now().Unix()
			//加密密码
			encryptedPasswordBytes, err2 := bcrypt.GenerateFromPassword([]byte(userPswd), bcrypt.DefaultCost)
			if err2 != nil {
				c.JSON(502, models.StringDataResponse{
					Status: 502,
					Data:   "",
					Msg:    "Having problem generating user info.",
					Error:  fmt.Sprintf("EncryptedPswd generation failed, error: %v", err2),
				})
				log.Printf("EncryptedPswd generation failed, error: %v", err2)
			} else {
				newUser.EncryptedPassword = string(encryptedPasswordBytes)
				//存储用户信息
				result := utils.DB.Create(&newUser)
				if result.Error != nil {
					c.JSON(502, models.StringDataResponse{
						Status: 502,
						Data:   "",
						Msg:    "Having problem saving user info.",
						Error:  fmt.Sprintf("Save user info failed, error: %v", result.Error),
					})
					log.Printf("Save user info failed, error: %v", result.Error)
				} else {
					c.JSON(200, models.StringDataResponse{
						Status: 200,
						Data:   "",
						Msg:    "Registration success!",
						Error:  "",
					})
					log.Printf("new user registration success! userId: %s, userName: %s, user info generation time: %d",
						newUser.Id, newUser.UserName, newUser.GenerateTime)
				}

			}
		}
	}
}

func Login(c *gin.Context) {
	userName := c.Query("userName")
	userPswd := c.Query("userPassword") //搜索用户

	var searchOutcome []models.User
	result := utils.DB.Where("user_name = ?", userName).First(&searchOutcome)

	// 登录逻辑： 搜索，比对，（返回错误/生成 token，返回 token）
	if fmt.Sprintf("%v", result.Error) == "record not found" {
		//无此用户：
		c.JSON(500, models.StringDataResponse{
			Status: 500,
			Data:   "",
			Msg:    "No such user.",
			Error:  fmt.Sprintf("No such user. userName: %s", userName),
		})
		log.Printf("No such user. name: %s.", userName)
	} else if result.Error != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem searching user in databases.",
			Error:  fmt.Sprintf("User search error: %v, userName: %s", result.Error, userName),
		})
		log.Printf("User search error: %v, userName: %s", result.Error, userName)
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(searchOutcome[0].EncryptedPassword), []byte(userPswd))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			//错误密码：
			c.JSON(500, models.StringDataResponse{
				Status: 500,
				Data:   "",
				Msg:    "Wrong password",
				Error:  "Wrong password",
			})
			log.Printf("Wrong password: user: %v", userName)
		} else if err == nil {

			//----------------------------------------------------

			// 签发 token
			token, err := utils.Sign(searchOutcome[0].Id, searchOutcome[0].UserName)
			if err != nil {
				showUnknownTokenError(c, err, userName, 1)
				return
			}
			c.JSON(200, models.JsonDataResponse{
				Status: 200,
				Data:   gin.H{"token": token},
				Msg:    "Login success",
				Error:  "Login success",
			})
			log.Printf("Login success: user: %v", userName)
		}
	}
}

func showUnknownTokenError(c *gin.Context, err error, userName string, stampPoint int) {
	c.JSON(502, models.StringDataResponse{
		Status: 502,
		Data:   "",
		Msg:    "Having problem giving token.",
		Error:  fmt.Sprintf("Having problem giving token. error: %v, userName: %s", err, userName),
	})
	log.Printf("Having problem giving token. error: %v, userName: %s, stampPoint: %d", err, userName, stampPoint)
}
