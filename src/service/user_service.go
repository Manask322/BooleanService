package service

import (
	"booleanservice/src/middleware"
	"booleanservice/src/models"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var user models.BooleanUser


func GetUser(c *gin.Context) models.BooleanUser{
	claims := jwt.ExtractClaims(c)
	var value models.User
	db := middleware.DB
	_ = db.Model(&models.User{}).Where("username = ?",claims["id"].(string)).Take(&value).Error

	return models.BooleanUser{
		Id : value.ID,
		UserName: claims["id"].(string),
		LastName: "LastName",
		FirstName: "FirstName",
	}
}


func SetUser(rUser models.BooleanUser)  {
	fmt.Println(rUser)
		user = rUser

}
