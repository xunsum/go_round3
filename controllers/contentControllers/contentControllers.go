package contentControllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go_test_project/models"
	"go_test_project/utils"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

func AddItem(c *gin.Context) {
	userId := c.PostForm("owner_id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	status, err1 := strconv.Atoi(c.PostForm("status"))
	endTime, err3 := strconv.Atoi(c.PostForm("end_time"))

	itemId, err2 := uuid.NewUUID()
	if err1 != nil || err2 != nil {
		//uuid 字符串转换错误：
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem generating user id.",
			Error:  fmt.Sprintf("Uuid generation failed error: strconv: %v, %v uuid: %v", err1, err3, err2),
		})
		log.Printf("Uuid generation failed error: strconv: %v, %v uuid: %v", err1, err3, err2)
		c.Abort()
	}

	newItem := models.TodoItem{
		OwnerId:    userId,
		Id:         itemId.String(),
		Title:      title,
		Content:    content,
		CreateTime: time.Now().Unix(),
		EndTime:    int64(endTime),
		Status:     status,
	}

	result := utils.DB.Create(&newItem)

	if result.Error != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem saving item info.",
			Error:  fmt.Sprintf("Save item info failed, error: %v", result.Error),
		})
		log.Printf("Save item  info failed, error: %v", result.Error)
	} else {
		c.JSON(200, models.StringDataResponse{
			Status: 200,
			Data:   "",
			Msg:    "Item save success!",
			Error:  "",
		})
		log.Printf("new item saving success! itemId: %s", newItem.Id)
	}
}

func SetItemState(c *gin.Context) {
	itemId := c.Query("item_id")
	newStateString := c.Query("new_state")
	newState, err := strconv.Atoi(newStateString)

	result := utils.DB.Model(&models.TodoItem{}).Where("id = ?", itemId).Update("status", newState)

	if result.Error != nil || err != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem changing item state.",
			Error:  fmt.Sprintf("Chenge item state failed, error: %v", result.Error),
		})
		log.Printf("Chenge item state failed, error: %v", result.Error)
	} else {
		c.JSON(200, models.StringDataResponse{
			Status: 200,
			Data:   "",
			Msg:    "Item state change success!",
			Error:  "",
		})
		log.Printf("item state change success! itemId: %s", itemId)
	}

}

func SetAllItemState(c *gin.Context) {
	userId := c.Query("user_id")
	newStateString := c.Query("new_state")
	newState, err := strconv.Atoi(newStateString)

	result := utils.DB.Model(&models.TodoItem{}).Where("owner_id = ?", userId).Update("status", newState)

	if result.Error != nil || err != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem changing all item states.",
			Error:  fmt.Sprintf("Chenge all item states failed, error: %v, %v", result.Error, err),
		})
		log.Printf("Chenge item state failed, error: %v", result.Error)
	} else {
		c.JSON(200, models.StringDataResponse{
			Status: 200,
			Data:   "",
			Msg:    "All item states change success!",
			Error:  "",
		})
		log.Printf("All item state change success! userId: %s", userId)
	}
}

func GetItems(c *gin.Context) {
	userId := c.Query("user_id")
	filterCondition := c.Query("filter_condition")
	var resultList []models.TodoItem
	var result *gorm.DB

	if filterCondition == "" {
		result = utils.DB.Model(&models.TodoItem{}).Where("owner_id = ?", userId).Find(&resultList)
	} else if filterCondition == "0" {
		result = utils.DB.Model(&models.TodoItem{}).Where("owner_id = ?", userId).Where("status = ?", 0).Find(&resultList)
	} else if filterCondition == "1" {
		result = utils.DB.Model(&models.TodoItem{}).Where("owner_id = ?", userId).Where("status = ?", 1).Find(&resultList)
	}

	if result.Error != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem getting items.",
			Error:  fmt.Sprintf("Get items failed, error: %v", result.Error),
		})
		log.Printf("Get items failed, error: %v", result.Error)
	} else {
		c.JSON(200, models.TodoItemResponse{
			Status: 200,
			Data: models.TodoData{
				Item:  resultList,
				Total: len(resultList),
			},
			Msg:   "Get items success!",
			Error: "",
		})
		log.Printf("Get items success! userId: %s", userId)
	}
}

func SearchItems(c *gin.Context) {
	userId := c.Query("user_id")
	searchContent := c.Query("search_for")
	var resultList []models.TodoItem
	var result *gorm.DB

	result = utils.DB.Model(&models.TodoItem{}).Where(utils.DB.Where("content LIKE ?", "%"+searchContent+"%").Or("title LIKE ?", "%"+searchContent+"%")).Find(&resultList)

	if result.Error != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem searching items.",
			Error:  fmt.Sprintf("Search items failed, error: %v", result.Error),
		})
		log.Printf("Search items failed, error: %v", result.Error)
	} else {
		c.JSON(200, models.TodoItemResponse{
			Status: 200,
			Data: models.TodoData{
				Item:  resultList,
				Total: len(resultList),
			},
			Msg:   "Search items success!",
			Error: "",
		})
		log.Printf("Search items success! userId: %s", userId)
	}
}

func DeleteItem(c *gin.Context) {
	itemId := c.Query("item_id")
	var result *gorm.DB

	result = utils.DB.Where("id = ?", itemId).Delete(models.TodoItem{})

	if result.Error != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem deleting item.",
			Error:  fmt.Sprintf("Delete item failed, error: %v", result.Error),
		})
		log.Printf("Search items failed, error: %v", result.Error)
	} else {
		c.JSON(200, models.StringDataResponse{
			Status: 200,
			Data:   "",
			Msg:    "Item delete success!",
			Error:  "",
		})
		log.Printf("Delete item success! itemId: %s", itemId)
	}
}

func DeleteItems(c *gin.Context) {
	userId := c.Query("user_id")
	filterCondition := c.Query("filter_condition")
	var result *gorm.DB

	if filterCondition == "" {
		result = utils.DB.Where("owner_id = ?", userId).Delete(models.TodoItem{})
	} else if filterCondition == "0" {
		result = utils.DB.Where("owner_id = ?", userId).Where("status = ?", 0).Delete(models.TodoItem{})
	} else if filterCondition == "1" {
		result = utils.DB.Where("owner_id = ?", userId).Where("status = ?", 1).Delete(models.TodoItem{})
	}

	if result.Error != nil {
		c.JSON(502, models.StringDataResponse{
			Status: 502,
			Data:   "",
			Msg:    "Having problem deleting items.",
			Error:  fmt.Sprintf("Delete items failed, error: %v", result.Error),
		})
		log.Printf("Search items failed, error: %v", result.Error)
	} else {
		c.JSON(200, models.StringDataResponse{
			Status: 200,
			Data:   "",
			Msg:    "Items delete success!",
			Error:  "",
		})
		log.Printf("Delete items success! userId: %s", userId)
	}
}
