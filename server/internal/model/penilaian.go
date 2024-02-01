package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Penilaian struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"type:uuid;primary_key"`
	AsesemenNonKognitif string    // not yet being used
	AsesemenKognitif    string
	AsesmenFormatif     string
	AsesmenSumatif      string
	Pengayaan           string
	Remedial            string
}
