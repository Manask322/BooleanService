package middleware

import (
	"booleanservice/src/models"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql" //COMMENT
	"github.com/jinzhu/gorm"
)

//DB is
var DB *gorm.DB
var err error
var Mu sync.RWMutex

func getDbURL() string {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	dbString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, database)
	return dbString
}

//StartDb is
func StartDb() (*gorm.DB, error) {

	DB, err = gorm.Open(os.Getenv("DB_DRIVER"), getDbURL())

	if err != nil {
		return nil, err
	}

	// var value models.NameValue
	DB.AutoMigrate(&models.NameValue{})
	DB.AutoMigrate(&models.User{})
	//DB.Model(&models.User{}).AddForeignKey("id", "name_values(user_id)", "CASCADE", "RESTRICT")

	//var user models.User
	//
	//user.Username = "manask322"
	//
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123Manas@"), 8)
	//user.Password = string(hashedPassword)
	//DB.Create(&user)

	return DB, err

}

//CloseDb iss
func CloseDb() {
	defer DB.Close()
}
