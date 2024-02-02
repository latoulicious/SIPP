package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Capaian struct {
	gorm.Model
	ID                         uuid.UUID `gorm:"type:uuid;primary_key"`
	JudulCapaian               string    `json:"judulCapaian"`
	UserID                     uuid.UUID `gorm:"type:uuid" json:"UserID"`
	User                       Users     `gorm:"foreignKey:UserID"`
	MapelID                    uuid.UUID `gorm:"type:uuid" json:"MapelID"`
	Mapel                      Mapel     `gorm:"foreignKey:MapelID"`
	KelasID                    uuid.UUID `gorm:"type:uuid" json:"KelasID"`
	Kelas                      Kelas     `gorm:"foreignKey:KelasID"`
	TahunAjarID                uuid.UUID `gorm:"type:uuid" json:"TahunAjarID"`
	TahunAjar                  TahunAjar `gorm:"foreignKey:TahunAjarID"`
	JudulElemen                string    `json:"judulElemen"`
	KetElemen                  string    `json:"ketElemen"`
	KetProsesMengamati         string    `json:"ketProsesMengamati"`
	KetProsesMempertanyakan    string    `json:"ketProsesMempertanyakan"`
	KetProsesMerencanakan      string    `json:"ketProsesMerencanakan"`
	KetProsesMemproses         string    `json:"ketProsesMemproses"`
	KetProsesMengevaluasi      string    `json:"ketProsesMengevaluasi"`
	KetProsesMengkomunikasikan string    `json:"ketProsesMengkomunikasikan"`
}
