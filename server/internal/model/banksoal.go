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
	Soal             string    `gorm:"type:varchar(255)"`
	OptionA          string    `gorm:"type:varchar(255)"`
	OptionB          string    `gorm:"type:varchar(255)"`
	OptionC          string    `gorm:"type:varchar(255)"`
	OptionD          string    `gorm:"type:varchar(255)"`
	OptionE          string    `gorm:"type:varchar(255)"`
	KunciJawaban     string    `gorm:"type:varchar(255)"`
	TingkatKesukaran string    `gorm:"type:varchar(255)"`
	IndikatorID      uuid.UUID `gorm:"type:uuid"`
	Indikator        Indikator `gorm:"foreignKey:IndikatorID"`
}
