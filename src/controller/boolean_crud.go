package controller

import (
	"booleanservice/src/middleware"
	"booleanservice/src/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var db *gorm.DB
var value models.NameValue

//NameValue is
type NameValue struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Value     bool
	Key       string
}

//VALUE is
type VALUE struct {
	Value *bool   `json:"value" binding:"required"`
	Key   *string `json:"key" binding:"required"`
}

func processRequest(c *gin.Context) (models.NameValue, error) {
	var rValue VALUE
	err := c.BindJSON(&rValue)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"err":  "Please submit the valid data",
		})
		return models.NameValue{}, err
	}
	bValue := models.NameValue{
		Value: *(rValue.Value),
		Key:   *(rValue.Key),
	}
	return bValue, nil
}

//CreateValue is
func CreateValue(c *gin.Context) {
	bValue, err := processRequest(c)

	if err != nil {
		return
	}

	queEle := middleware.ServiceQueueElement{C: c}
	if len(middleware.ServiceQueueIn) > middleware.SERVICEQUEUELENGTH {
		queEle.C.JSON(500, gin.H{
			"msg": "Database Server Overload!!!! Please Try Again Later",
		})
		return
	}
	middleware.ServiceQueueIn <- queEle
	log.Printf("| Length of Create Process Queue : %d", len(middleware.ServiceQueueIn))
	middleware.StartQueueJob(c, bValue)
	queEle = <-middleware.ServiceQueueOut

	queEle.C.JSON(201, gin.H{
		"id":    queEle.NameValue.ID,
		"key":   queEle.NameValue.Key,
		"value": queEle.NameValue.Value,
	})
	return
}

//UpdateValue is
func UpdateValue(c *gin.Context) {
	id := c.Param("id")

	bValue, err := processRequest(c)
	if err != nil {
		return
	}

	db = middleware.DB
	var nameValue models.NameValue
	err = db.Model(&nameValue).Where("id = ?", id).Update(
		map[string]interface{}{"key": bValue.Key, "value": bValue.Value}).Error

	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = db.Model(&models.NameValue{}).Where("id = ?", id).Take(&nameValue).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(400, gin.H{
			"err": "Record Not found",
		})
		return
	}
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"id":    nameValue.ID,
		"key":   nameValue.Key,
		"value": nameValue.Value,
	})
	return

}

//DeleteValue is
func DeleteValue(c *gin.Context) {
	id := c.Param("id")
	db = middleware.DB
	err := db.Model(&models.NameValue{}).Where("id = ?", id).Take(&value).Error
	if err != nil {
		c.JSON(400, gin.H{
			"err": "Key Not found",
		})
		return
	}
	err = db.Unscoped().Where("id = ?", id).Delete(&models.NameValue{}).Error

	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(204, gin.H{
		"message": "No Content",
	})

}

//ListAll is
func ListAll(c *gin.Context) {
	db = middleware.DB
	c.JSON(200, gin.H{
		"message": "Boolean Service",
	})
}

//GetValue is
func GetValue(c *gin.Context) {
	id := c.Param("id")
	db = middleware.DB
	err := db.Model(&models.NameValue{}).Where("id = ?", id).Take(&value).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(400, gin.H{
			"err": "Record Not found",
		})
		return
	}
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"id":    value.ID,
		"key":   value.Key,
		"value": value.Value,
	})
	return
}
