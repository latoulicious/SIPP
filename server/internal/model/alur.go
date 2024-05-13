package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlurTP struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;primary_key"`
	Elemen             string    `gorm:"type:varchar(255)"`
	LingkupMateri      string    `gorm:"type:text"`
	TujuanPembelajaran string    `gorm:"type:text"`
	KodeTP             string    `gorm:"type:varchar(30)"`
	AlokasiWaktu       string    `gorm:"type:varchar(30)"`
	SumberBelajar      string    `gorm:"type:text"`
	ProjekPPancasila   string    `gorm:"type:text"`
}
