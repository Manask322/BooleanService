package repo

import (
	"booleanservice/src/middleware"
	"booleanservice/src/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

func CreateRecord(bValue models.NameValue) (*gorm.DB) {
	db := middleware.DB
	response := db.Create(&bValue)
	return response
}

func UpdateRecord(id string,bValue models.NameValue,userId int )(models.NameValue,models.BooleanError) {
	db := middleware.DB
	var nameValue models.NameValue
	err, nameValue := GetRecordWithUserId(id,userId)
	if err.Code != 0 {
		return models.NameValue{},err
	}
	updateErr := db.Model(&models.NameValue{}).Where("id = ?", id).Update(
		map[string]interface{}{"key": bValue.Key, "value": bValue.Value}).Error

	if updateErr != nil {
		return models.NameValue{},models.BooleanError{Code: 4,Status: 500}
	}
	middleware.Mu.RLock()
	err, nameValue = GetRecordWithUserId(id,userId)
	if err.Code != 0 {
		return models.NameValue{},err
	}
	return nameValue,models.BooleanError{Status: 200}
}

func DeleteRecord(id string,userId int) (models.NameValue,models.BooleanError){
	db := middleware.DB
	err, _ := GetRecordWithUserId(id,userId)
	if err.Code != 0 {
		return models.NameValue{},err
	}

	dErr := db.Unscoped().Where("id = ?", id).Delete(&models.NameValue{}).Error
	fmt.Println(" DBError : ", dErr)
	if dErr != nil {
		return models.NameValue{},models.BooleanError{Code:5,Status: 500}
	}
	return models.NameValue{}, models.BooleanError{Status: 204}
}

func GetRecord(id string,userId int) (models.NameValue,models.BooleanError) {
	err, value := GetRecordWithUserId(id,userId)
	if err.Code != 0 {
		return models.NameValue{},err
	}
	return value,err
}

func GetRecordWithUserId(id string,userId int) (models.BooleanError,models.NameValue){
	db := middleware.DB
	var value models.NameValue
	middleware.Mu.RLock()
	err := db.Model(&models.NameValue{}).Where("id = ? AND user_id = ?", id,userId).Take(&value).Error
	middleware.Mu.RUnlock()
	if gorm.IsRecordNotFoundError(err){
		return models.BooleanError{Code:3,Status: 400}, models.NameValue{}
	}
	if err != nil {
		return models.BooleanError{Code:3,Status: 500}, models.NameValue{}
	}
	return models.BooleanError{Status: 200}, value
}