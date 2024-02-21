package model

import "github.com/google/uuid"

type DynamicField struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type CreateSoalItem struct {
	BankSoalID    uuid.UUID      `json:"bankSoalId"`
	DynamicFields []DynamicField `json:"dynamicFields"`
}

type CreateSoalRequest struct {
	TipeSoal  string           `json:"tipeSoal"`
	Hari      string           `json:"hari"`
	Tanggal   string           `json:"tanggal"`
	Waktu     string           `json:"waktu"`
	UserID    uuid.UUID        `json:"userId"`
	MapelID   uuid.UUID        `json:"mapelId"`
	KelasID   uuid.UUID        `json:"kelasId"`
	JurusanID uuid.UUID        `json:"jurusanId"`
	Items     []CreateSoalItem `json:"items"`
	//  Add any other Soal related fields directly here
}
