package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankSoal struct {
	gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID           uuid.UUID `gorm:"type:uuid" json:"UserID"`
	User             Users     `gorm:"foreignKey:UserID"`
	MapelID          uuid.UUID `gorm:"type:uuid" json:"MapelID"`
	Mapel            Mapel     `gorm:"foreignKey:MapelID"`
	KelasID          uuid.UUID `gorm:"type:uuid" json:"KelasID"`
	Kelas            Kelas     `gorm:"foreignKey:KelasID"`
	JurusanID        uuid.UUID `gorm:"type:uuid" json:"JurusanID"`
	Jurusan          Jurusan   `gorm:"foreignKey:JurusanID"`
	TahunAjarID      uuid.UUID `gorm:"type:uuid" json:"TahunAjarID"`
	TahunAjar        TahunAjar `gorm:"foreignKey:TahunAjarID"`
	Soal             string
	OptionA          string
	OptionB          string
	OptionC          string
	OptionD          string
	OptionE          string
	KunciJawaban     string
	Materi           string
	Indikator        string
	TingkatKesukaran string
}
