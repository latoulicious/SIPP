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
	Sekolah              string    `gorm:"type:varchar(50)" json:"sekolah"`
	AlokasiWaktu         string    `gorm:"type:varchar(30)" json:"alokasiWaktu"`
	KompetensiAwal       string    `gorm:"type:text" json:"kompetensiAwal"`
	ProjekPPancasila     string    `gorm:"type:text" json:"projekPPancasila"`
	SaranaPrasarana      string    `gorm:"type:text" json:"saranaPrasarana"`
	TargetPesertaDidik   string    `gorm:"type:text" json:"targetPesertaDidik"`
	ModelPembelajaran    string    `gorm:"type:text" json:"modelPembelajaran"`
	TujuanPembelajaran   string    `gorm:"type:text" json:"tujuanPembelajaran"`
	PemahamanBermakna    string    `gorm:"type:text" json:"pemahamanBermakna"`
	PertanyaanPemantik   string    `gorm:"type:text" json:"pertanyaanPemantik"`
	KegiatanPembelajaran string    `gorm:"type:text" json:"kegiatanPembelajaran"`
	Asesnmen             string    // not yet being used
	PengayaanRemedial    string    // not yet being used
	Refleksi             string    `gorm:"type:text" json:"refleksi"`
	Glosarium            string    `gorm:"type:text" json:"glosarium"`
	DaftarPustaka        string    `gorm:"type:text" json:"daftarPustaka"`
}
