package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Capaian struct {
	gorm.Model
	ID                         uuid.UUID `gorm:"type:uuid;primary_key"`
	JudulCapaian               string
	User                       Users     `gorm:"foreignkey:UserID"`
	UserID                     uuid.UUID `gorm:"type:uuid"`
	Mapel                      Mapel     `gorm:"foreignkey:MapelID"`
	MapelID                    uuid.UUID `gorm:"type:uuid"`
	Kelas                      Kelas     `gorm:"foreignkey:KelasID"`
	KelasID                    uuid.UUID `gorm:"type:uuid"`
	TahunAjar                  TahunAjar `gorm:"foreignkey:TahunAjarID"`
	TahunAjarID                uuid.UUID `gorm:"type:uuid"`
	JudulElemen                string
	KetElemen                  string
	KetProsesMengamati         string
	KetProsesMempertanyakan    string
	KetProsesMerencanakan      string
	KetProsesMemproses         string
	KetProsesMengevaluasi      string
	KetProsesMengkomunikasikan string
}
