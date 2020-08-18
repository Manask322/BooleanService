package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//NameValue is
type NameValue struct {
	ID        uuid.UUID `gorm:"primary_key; type:char(36); column:id;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Value     bool      `gorm:"not null"`
	Key       string
}

//BeforeCreate is
func (base *NameValue) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
