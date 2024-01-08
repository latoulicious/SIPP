package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlurTP struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;primary_key"`
	Elemen             string
	LingkupMateri      string
	TujuanPembelajaran string
	KodeTP             string
	AlokasiWaktu       string
	SumberBelajar      string
	ProjekPPancasila   string
}
