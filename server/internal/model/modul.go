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
	Sekolah              string    `json:"sekolah"`
	AlokasiWaktu         string    `json:"alokasiWaktu"`
	KompetensiAwal       string    `json:"kompetensiAwal"`
	ProjekPPancasila     string    `json:"projekPPancasila"`
	SaranaPrasarana      string    `json:"saranaPrasarana"`
	TargetPesertaDidik   string    `json:"targetPesertaDidik"`
	ModelPembelajaran    string    `json:"modelPembelajaran"`
	TujuanPembelajaran   string    `json:"tujuanPembelajaran"`
	PemahamanBermakna    string    `json:"pemahamanBermakna"`
	PertanyaanPemantik   string    `json:"pertanyaanPemantik"`
	KegiatanPembelajaran string    `json:"kegiatanPembelajaran"`
	Asesnmen             string    // not yet being used
	PengayaanRemedial    string    // not yet being used
	Refleksi             string    `json:"refleksi"`
	Glosarium            string    `json:"glosarium"`
	DaftarPustaka        string    `json:"daftarPustaka"`
}
