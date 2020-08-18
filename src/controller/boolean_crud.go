package controller

import (
	"booleanservice/src/middleware"
	"booleanservice/src/models"
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Println("Inside Main Func")
	bValue, err := processRequest(c)

	if err != nil {
		return
	}

	db = middleware.DB
	response := db.Create(&bValue)

	p, _ := json.Marshal(response.Value)
	var nameValue NameValue
	err = json.Unmarshal(p, &nameValue)

	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    nameValue.ID,
		"key":   nameValue.Key,
		"value": nameValue.Value,
	})
	return
}

//UpdateValue is
func UpdateValue(c *gin.Context) {
	id := c.Param("id")

	bValue, err := processRequest(c)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}

	db = middleware.DB
	var nameValue models.NameValue
	err = db.Model(&nameValue).Where("id = ?", id).Update(
		map[string]interface{}{"key": bValue.Key, "value": bValue.Value}).Error

	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
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
