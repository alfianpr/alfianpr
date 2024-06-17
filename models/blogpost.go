package models

import (
	"time"

	"gorm.io/gorm"
)

type BlogPost struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	HTMLContent string    `json:"html_content"`
	CreatedAt   time.Time `json:"created_at"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&BlogPost{})
}
