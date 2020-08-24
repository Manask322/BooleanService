package service

import (
	"booleanservice/src/middleware"
	"booleanservice/src/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

// User demo
type User struct {
	Id 		  int
	UserName  string
	FirstName string
	LastName  string
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//JWTMiddleware is
func JWTMiddleware() (*jwt.GinJWTMiddleware, error) {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				fmt.Println("V : ", v)
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			fmt.Println("claims : ", claims)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					Id : 0,
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}
			fmt.Println(userID, " : ", password)
			db := middleware.DB
			var value models.User
			err := db.Model(&models.User{}).Where("username = ?",userID).Take(&value).Error
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			err = bcrypt.CompareHashAndPassword([]byte(value.Password), []byte(password))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication

			}
			user := models.BooleanUser{
				Id : value.ID,
				UserName: userID,
				LastName: "LastName",
				FirstName: "FirstName",
			}
			fmt.Println("User Set")
			SetUser(user)
			return &User{
				Id : value.ID,
				UserName: userID,
				LastName: "LastName",
				FirstName: "FirstName",
			}, nil

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)
			fmt.Println("ID : ", claims["id"])
			if _, ok := data.(*User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil, err
	}
	return authMiddleware, nil
}
