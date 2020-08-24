package service

import (
	"booleanservice/src/models"
	"booleanservice/src/repo"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gofrs/uuid"
)

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
	Value *bool `json:"value" binding:"required"`
}

type KEY struct {
	Key *string `json:"key" binding:"required"`
}

//KeyValue is
type KeyValue struct {
	Value *bool   `json:"value" binding:"required"`
	Key   *string `json:"key" binding:"required"`
}

// BooleanCreateService is
func BooleanCreateService(user models.BooleanUser,bValue models.NameValue) (models.DatabaseNameValue,models.BooleanError){

	if user.Id == 0 {
		return models.DatabaseNameValue{},models.BooleanError{Code: 1,Status: 401}
	}

	//queEle := middleware.ServiceQueueElement{C: c}
	//if len(middleware.ServiceQueueIn) > middleware.SERVICEQUEUELENGTH {
	//	queEle.C.JSON(500, gin.H{
	//		"msg": "Database Server Overload!!!! Please Try Again Later",
	//	})
	//	return middleware.ServiceQueueElement{},err
	//}
	//middleware.ServiceQueueIn <- queEle
	//log.Printf("| Length of Create Process Queue : %d", len(middleware.ServiceQueueIn))
	bValue.UserID = user.Id
	response := repo.CreateRecord(bValue)
	p, _ := json.Marshal(response.Value)
	var nameValue models.DatabaseNameValue
	json.Unmarshal(p, &nameValue)
	//middleware.StartQueueJob(c, bValue)


	return nameValue,models.BooleanError{Status: 201}
}

func BooleanUpdateService(id string,user models.BooleanUser,bValue models.NameValue) (models.NameValue,models.BooleanError){
	if user.Id == 0 {
		return models.NameValue{},models.BooleanError{Code: 1,Status: 401}
	}


	 return repo.UpdateRecord(id,bValue,user.Id)
}

func BooleanGetService(id string,userId int) (models.NameValue,models.BooleanError){
	if userId == 0 {
		return models.NameValue{},models.BooleanError{Code: 1,Status: 401}
	}
	return repo.GetRecord(id,userId)
}

func BooleanDeleteService(id string,userId int) (models.NameValue,models.BooleanError){
	if userId == 0 {
		return models.NameValue{},models.BooleanError{Code: 1,Status: 401}
	}
	return repo.DeleteRecord(id,userId)
}


/*
	HELPER FUNCTIONS
*/


func ProcessRequest(c *gin.Context) (models.NameValue, models.BooleanError) {
	var rValue VALUE
	var kValue KeyValue
	berr := c.ShouldBindBodyWith(&rValue, binding.JSON);
	aerr := c.ShouldBindBodyWith(&kValue, binding.JSON);
	user = GetUser(c)
	if berr != nil {
		return models.NameValue{}, models.BooleanError{Code: 2, Status: 500}
	}
	if  aerr == nil {
		bValue := models.NameValue{
			Value: *(kValue.Value),
			Key:   *(kValue.Key),
		}
		return bValue, models.BooleanError{Status: 200}
	} else if  berr == nil {
		bValue := models.NameValue{
			Value: *(rValue.Value),
			Key:   "",
		}
		return bValue, models.BooleanError{Status: 200}
	}

	return models.NameValue{}, models.BooleanError{Code: 4, Status: 500}
}
