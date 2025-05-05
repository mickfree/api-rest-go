package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Album struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title      string    `json:"title"`
	Artist     string    `json:"artist"`
	Year       int       `json:"year"`
	Genre      string    `json:"genre"`
	Language   string    `json:"language"`
	Duration   int64     `json:"duration"`
	CoverImage string    `json:"coverimage"`
}

// add UUID
func (album *Album) BeforeCreate(tx *gorm.DB) (err error) {
	album.ID = uuid.New()
	return
}
