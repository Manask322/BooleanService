package middleware

import (
	"booleanservice/src/models"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

const SERVICEQUEUELENGTH = 10

//ServiceQueueIn is a Buffered Channel to store the incoming requests
var ServiceQueueIn = make(chan ServiceQueueElement, SERVICEQUEUELENGTH)

var ServiceQueueOut = make(chan ServiceQueueElement, SERVICEQUEUELENGTH)

type ServiceQueueElement struct {
	C         *gin.Context
	NameValue models.DatabaseNameValue
}

//StartQueueJob is
func StartQueueJob(c *gin.Context, bValue models.NameValue) {

	go processDatabase(bValue)
}

func processDatabase(bValue models.NameValue) {
	db := DB
	response := db.Create(&bValue)
	p, _ := json.Marshal(response.Value)
	var nameValue models.DatabaseNameValue
	json.Unmarshal(p, &nameValue)
	queEle := <-ServiceQueueIn
	queEle.NameValue = nameValue
	ServiceQueueOut <- queEle
}
