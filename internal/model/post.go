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

var ExamplePost = Post{
	Title:     "Introduction to Golang",
	Content:   "Golang, or Go, is an open-source programming language designed for simplicity and efficiency...",
	Category:  "Programming",
	Tags:      []string{"Golang", "Programming", "Tutorial"},
	UserID:    1, // Assume this user ID exists in the database
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
