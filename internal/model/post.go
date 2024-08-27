package model

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	Category  string    `gorm:"not null" json:"category"`
	Tags      []string  `gorm:"type:text[]" json:"tags"`
	UserID    uint      `gorm:"not null" json:"userID"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
