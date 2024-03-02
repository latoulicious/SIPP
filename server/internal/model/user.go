// internal/model/user.go

package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Username string    `gorm:"type:varchar(20)"`
	Password string    `gorm:"type:varchar(60)"`
	Name     string    `gorm:"type:varchar(50)"`
	NIP      string    `gorm:"column:NIP;type:varchar(20)"` // Specify the column name here
	Golongan string    `gorm:"type:varchar(20)"`
	Role     string    `gorm:"type:varchar(20)"`
}
