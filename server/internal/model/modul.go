package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModulAjar struct {
	gorm.Model
	ID                   uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID               uuid.UUID `gorm:"type:uuid" json:"UserID"`
	User                 Users     `gorm:"foreignKey:UserID"`
	TahunAjarID          uuid.UUID `gorm:"type:uuid" json:"TahunAjarID"`
	TahunAjar            TahunAjar `gorm:"foreignKey:TahunAjarID"`
	KelasID              uuid.UUID `gorm:"type:uuid" json:"KelasID"`
	Kelas                Kelas     `gorm:"foreignKey:KelasID"`
	Sekolah              string    `gorm:"type:varchar(255)" json:"sekolah"`
	AlokasiWaktu         string    `gorm:"type:varchar(255)" json:"alokasiWaktu"`
	KompetensiAwal       string    `gorm:"type:varchar(255)" json:"kompetensiAwal"`
	ProjekPPancasila     string    `gorm:"type:varchar(255)" json:"projekPPancasila"`
	SaranaPrasarana      string    `gorm:"type:varchar(255)" json:"saranaPrasarana"`
	TargetPesertaDidik   string    `gorm:"type:varchar(255)" json:"targetPesertaDidik"`
	ModelPembelajaran    string    `gorm:"type:varchar(255)" json:"modelPembelajaran"`
	TujuanPembelajaran   string    `gorm:"type:varchar(255)" json:"tujuanPembelajaran"`
	PemahamanBermakna    string    `gorm:"type:varchar(255)" json:"pemahamanBermakna"`
	PertanyaanPemantik   string    `gorm:"type:varchar(255)" json:"pertanyaanPemantik"`
	KegiatanPembelajaran string    `gorm:"type:varchar(255)" json:"kegiatanPembelajaran"`
	Asesnmen             string    // not yet being used
	PengayaanRemedial    string    // not yet being used
	Refleksi             string    `gorm:"type:varchar(255)" json:"refleksi"`
	Glosarium            string    `gorm:"type:varchar(255)" json:"glosarium"`
	DaftarPustaka        string    `gorm:"type:varchar(255)" json:"daftarPustaka"`
}
