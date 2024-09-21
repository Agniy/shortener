package models

import "time"

type Links struct {
	ID        uint64 `gorm:"primaryKey"`
	URL       string
	ShortURL  string `gorm:"index"`
	CreatedAt time.Time
}
