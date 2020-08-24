package controller

import (
	"booleanservice/src/models"
	"booleanservice/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

var err error

//CreateValue is
func CreateValue(c *gin.Context) {
	bValue, err := service.ProcessRequest(c)
	if err.Code != 0 {
		c.JSON(400, gin.H{
			"message": models.GetErrorMessage(err.Code),
		})
		return
	}
	user := service.GetUser(c)
	nameValue, err := service.BooleanCreateService(user,bValue)
	var code int
	if err.Status == 401 {
		code = 401
	}else{
		code = 400
	}
	if err.Code != 0 {
		c.JSON(code, gin.H{
			"message": models.GetErrorMessage(err.Code),
		})
		return
	}
	//queEle = <-middleware.ServiceQueueOut
	c.JSON(201, gin.H{
		"id":    nameValue.ID,
		"key":   nameValue.Key,
		"value": nameValue.Value,
		"User": user.UserName,
	})

}

//UpdateValue is
func UpdateValue(c *gin.Context) {
	id := c.Param("id")
	bValue, err := service.ProcessRequest(c)
	if err.Code != 0 {
		c.JSON(400, gin.H{
			"message": models.GetErrorMessage(err.Code),
		})
		return
	}
	user := service.GetUser(c)
	nameValue, err := service.BooleanUpdateService(id,user,bValue)
	if err.Code != 0 {
		c.JSON(400, gin.H{
			"message": models.GetErrorMessage(err.Code),
		})
		return
	}
	//queEle = <-middleware.ServiceQueueOut
	c.JSON(201, gin.H{
		"id":    nameValue.ID,
		"key":   nameValue.Key,
		"value": nameValue.Value,
	})
}

//DeleteValue is
func DeleteValue(c *gin.Context) {
	id := c.Param("id")
	user := service.GetUser(c)
	_,err := service.BooleanDeleteService(id,user.Id)
	fmt.Println(" C : ", err)
	if err.Code != 0 {
		c.JSON(err.Status, gin.H{
			"message": models.GetErrorMessage(err.Code),
		})
		return
	}
	c.AbortWithStatus(err.Status)
	return

}


//GetValue is
func GetValue(c *gin.Context) {
	id := c.Param("id")
	user := service.GetUser(c)
	value, err := service.BooleanGetService(id,user.Id)
	if err.Code != 0 {
		c.JSON(400, gin.H{
			"message": models.GetErrorMessage(err.Code),
		})
		return
	}
	c.JSON(err.Status, gin.H{
		"id":    value.ID,
		"key":   value.Key,
		"value": value.Value,
	})
}