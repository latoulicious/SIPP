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
	Soal             string    `gorm:"type:text"`
	OptionA          string    `gorm:"type:text"`
	OptionB          string    `gorm:"type:text"`
	OptionC          string    `gorm:"type:text"`
	OptionD          string    `gorm:"type:text"`
	OptionE          string    `gorm:"type:text"`
	KunciJawaban     string    `gorm:"type:text"`
	Materi           string    `gorm:"type:text"`
	Indikator        string    `gorm:"type:text"`
	TingkatKesukaran string    `gorm:"type:text"`
}
