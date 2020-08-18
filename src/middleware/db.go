package middleware

import (
	"booleanservice/src/models"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //COMMENT
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

//DB is
var DB *gorm.DB
var err error

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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB, err = gorm.Open(os.Getenv("DB_DRIVER"), getDbURL())

	if err != nil {
		return nil, err
	}

	// var value models.NameValue
	DB.AutoMigrate(&models.NameValue{})
	return DB, err

}

//CloseDb iss
func CloseDb() {
	defer DB.Close()
}
