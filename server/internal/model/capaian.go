package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Capaian struct {
	gorm.Model
	ID                         uuid.UUID `gorm:"type:uuid;primary_key"`
	JudulCapaian               string    `gorm:"type:varchar(255)" json:"judulCapaian"`
	UserID                     uuid.UUID `gorm:"type:uuid" json:"UserID"`
	User                       Users     `gorm:"foreignKey:UserID"`
	MapelID                    uuid.UUID `gorm:"type:uuid" json:"MapelID"`
	Mapel                      Mapel     `gorm:"foreignKey:MapelID"`
	KelasID                    uuid.UUID `gorm:"type:uuid" json:"KelasID"`
	Kelas                      Kelas     `gorm:"foreignKey:KelasID"`
	TahunAjarID                uuid.UUID `gorm:"type:uuid" json:"TahunAjarID"`
	TahunAjar                  TahunAjar `gorm:"foreignKey:TahunAjarID"`
	JudulElemen                string    `gorm:"type:varchar(255)" json:"judulElemen"`
	KetElemen                  string    `gorm:"type:text" json:"ketElemen"`
	KetProsesMengamati         string    `gorm:"type:text" json:"ketProsesMengamati"`
	KetProsesMempertanyakan    string    `gorm:"type:text" json:"ketProsesMempertanyakan"`
	KetProsesMerencanakan      string    `gorm:"type:text" json:"ketProsesMerencanakan"`
	KetProsesMemproses         string    `gorm:"type:text" json:"ketProsesMemproses"`
	KetProsesMengevaluasi      string    `gorm:"type:text" json:"ketProsesMengevaluasi"`
	KetProsesMengkomunikasikan string    `gorm:"type:text" json:"ketProsesMengkomunikasikan"`
}
