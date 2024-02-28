package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlurTP struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;primary_key"`
	Elemen             string    `gorm:"type:varchar(255)"`
	LingkupMateri      string    `gorm:"type:varchar(255)"`
	TujuanPembelajaran string    `gorm:"type:varchar(255)"`
	KodeTP             string    `gorm:"type:varchar(255)"`
	AlokasiWaktu       string    `gorm:"type:varchar(255)"`
	SumberBelajar      string    `gorm:"type:varchar(255)"`
	ProjekPPancasila   string    `gorm:"type:varchar(255)"`
}
