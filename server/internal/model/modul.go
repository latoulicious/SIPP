package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModulAjar struct {
	gorm.Model
	ID                   uuid.UUID `gorm:"type:uuid;primary_key"`
	User                 Users     `gorm:"foreignkey:UserID"` // Foreign key to Users table
	UserID               uuid.UUID `gorm:"type:uuid"`
	TahunAjar            TahunAjar `gorm:"foreignkey:TahunAjarID"` // Foreign key to TahunAjar table
	TahunAjarID          uuid.UUID `gorm:"type:uuid"`
	Kelas                Kelas     `gorm:"foreignkey:KelasID"` // Foreign key to Kelas table
	KelasID              uuid.UUID `gorm:"type:uuid"`
	Sekolah              string
	AlokasiWaktu         string
	KompetensiAwal       string
	ProjekPPancasila     string
	SaranaPrasarana      string
	TargetPesertaDidik   string
	ModelPembelajaran    string
	TujuanPembelajaran   string
	PemahamanBermakna    string
	PertanyaanPemantik   string
	KegiatanPembelajaran string
	Asesnmen             string
	PengayaanRemedial    string
	Refleksi             string
	Glosarium            string
	DaftarPustaka        string
}
