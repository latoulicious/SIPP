package model

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JSON json.RawMessage

// UnmarshalJSON implements the json.Unmarshaler interface.
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append(*j, data...)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

type Formatif struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID `gorm:"type:uuid" json:"UserID"`
	User          Users     `gorm:"foreignKey:UserID"`
	MapelID       uuid.UUID `gorm:"type:uuid" json:"MapelID"`
	Mapel         Mapel     `gorm:"foreignKey:MapelID"`
	KelasID       uuid.UUID `gorm:"type:uuid" json:"KelasID"`
	Kelas         Kelas     `gorm:"foreignKey:KelasID"`
	TahunAjarID   uuid.UUID `gorm:"type:uuid" json:"TahunAjarID"`
	TahunAjar     TahunAjar `gorm:"foreignKey:TahunAjarID"`
	Pertanyaan    string    `gorm:"type:varchar(255)"`
	DynamicFields JSON      `gorm:"type:jsonb" json:"DynamicFields"` // Corrected field definition
}
