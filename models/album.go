package models

type Album struct {
	ID       string `gorm:"primaryKey"`
	Title    string
	Artist   string
	Year     int
	Genre    string
	Language string
	Duration int64
}
