package repo_test

import (
	"booleanservice/src/middleware"
	"booleanservice/src/models"
	"booleanservice/src/repo"
	"encoding/json"
	"testing"
)

func TestGet(t *testing.T) {
	_, err1 := middleware.StartDb()
	if err1 != nil {
		t.Fatalf("Error :  %s", err1)
	}

	_, err := repo.GetRecord("32323232", 1)
	if err.Code != 0 {
		t.Logf("Error :  %s", models.GetErrorMessage(err.Code))
	}
}

func TestCreateRecord(t *testing.T) {
	_, err1 := middleware.StartDb()
	if err1 != nil {
		t.Fatalf("Error :  %s", err1)
	}
	bValue := models.NameValue{
		Value:  true,
		Key:    "Testing Create",
		UserID: 1,
	}
	response := repo.CreateRecord(bValue)
	p, _ := json.Marshal(response.Value)
	var nameValue models.DatabaseNameValue
	json.Unmarshal(p, &nameValue)
	if nameValue.Key != bValue.Key {
		t.Fatalf("Creation Error")
	}
	_, err := repo.GetRecord(nameValue.ID.String(), bValue.UserID)
	if err.Code != 0 {
		t.Fatalf("Error :  %s", models.GetErrorMessage(err.Code))
	}

	bValue.Value = false
	_, err = repo.UpdateRecord(nameValue.ID.String(),bValue,bValue.UserID)
	if err.Code != 0 {
		t.Fatalf("Error : %s", models.GetErrorMessage(err.Code))
	}

	_ ,err = repo.DeleteRecord(nameValue.ID.String(),bValue.UserID)
	if err.Code != 0 {
		t.Fatalf("Error : %s", models.GetErrorMessage(err.Code))
	}
	_ ,err = repo.DeleteRecord(nameValue.ID.String(),bValue.UserID)
	if err.Code == 0 {
		t.Fatalf("Error : Record Should Not be found.")
	}

}
