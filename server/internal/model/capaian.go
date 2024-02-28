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
	JudulElemen                string    `gorm:"type:varchar(255)" json:"judulElemen"`
	KetElemen                  string    `gorm:"type:varchar(255)" json:"ketElemen"`
	KetProsesMengamati         string    `gorm:"type:varchar(255)" json:"ketProsesMengamati"`
	KetProsesMempertanyakan    string    `gorm:"type:varchar(255)" json:"ketProsesMempertanyakan"`
	KetProsesMerencanakan      string    `gorm:"type:varchar(255)" json:"ketProsesMerencanakan"`
	KetProsesMemproses         string    `gorm:"type:varchar(255)" json:"ketProsesMemproses"`
	KetProsesMengevaluasi      string    `gorm:"type:varchar(255)" json:"ketProsesMengevaluasi"`
	KetProsesMengkomunikasikan string    `gorm:"type:varchar(255)" json:"ketProsesMengkomunikasikan"`
}
