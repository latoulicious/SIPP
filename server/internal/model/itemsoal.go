package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemSoal struct {
	gorm.Model
	ID            uuid.UUID     `gorm:"type:uuid;primary_key"`
	SoalID        uuid.UUID     `gorm:"type:uuid"`
	Soal          *Soal         `gorm:"foreignKey:SoalID;associationForeignKey:SoalID"`
	BankSoalID    uuid.UUID     `gorm:"type:uuid"`
	BankSoal      BankSoal      `gorm:"foreignKey:BankSoalID"`
	DynamicFields DynamicFields `gorm:"type:jsonb" json:"DynamicFields"`
	QuestionCount int64         `json:"questionCount"`
}

type DynamicFields map[string]interface{}

// Implementing the Scanner interface for DynamicFields
func (df *DynamicFields) Scan(value interface{}) error {
	if value == nil {
		*df = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	var v map[string]interface{}
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}
	*df = v
	return nil
}

// Implementing the Valuer interface for DynamicFields
func (df DynamicFields) Value() (driver.Value, error) {
	if df == nil {
		return nil, nil
	}
	return json.Marshal(df)
}
