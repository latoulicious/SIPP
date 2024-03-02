package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlurTP struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;primary_key"`
	Elemen             string    `gorm:"type:text"`
	LingkupMateri      string    `gorm:"type:text"`
	TujuanPembelajaran string    `gorm:"type:text"`
	KodeTP             string    `gorm:"type:text"`
	AlokasiWaktu       string    `gorm:"type:text"`
	SumberBelajar      string    `gorm:"type:text"`
	ProjekPPancasila   string    `gorm:"type:text"`
}
