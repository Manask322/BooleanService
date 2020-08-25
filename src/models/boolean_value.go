package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       int    `gorm:"primary key"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type BooleanUser struct {
	Id        int
	UserName  string
	FirstName string
	LastName  string
}

//NameValue is
type NameValue struct {
	ID        uuid.UUID `gorm:"primary_key; type:char(36); column:id;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Value     bool      `gorm:"not null"`
	Key       string
	UserID    int `json:"user_id"`
}

//BeforeCreate is
func (base *NameValue) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

//DatabaseNameValue is
type DatabaseNameValue struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Value     bool
	Key       string
	UserID    int
}
