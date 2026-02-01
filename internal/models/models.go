package models

import "time"

type Chat struct {
	ID        int        `gorm:"primaryKey"`
	Title     string     `gorm:"not null;size:200"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	Messages  []Messages `gorm:"constraint:OnDelete:CASCADE"`
}

type Messages struct {
	ID        int       `gorm:"primaryKey"`
	ChatID    int       `gorm:"index;not null"`
	Text      string    `gorm:"not null;size:5000"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
