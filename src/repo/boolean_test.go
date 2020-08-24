package repo_test

import (
	"booleanservice/src/middleware"
	"booleanservice/src/models"
	"booleanservice/src/repo"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	middleware.StartDb()
	err,w := repo.GetRecord("32323232", 1)
	fmt.Println(w)
	if err.Code != 0 {
		t.Fatalf("Error :  %s", models.GetErrorMessage(err.Code))
	}
}
