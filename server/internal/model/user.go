// internal/model/users.go

package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Username string
	Password string
	Name     string
	Mapel    string
	Role     string
}
