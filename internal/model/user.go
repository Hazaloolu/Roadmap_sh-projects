package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique; not null" json:"username"`
	Email     string    `gorm:"unique; not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Posts     []Post    `gorm:"foreignKey:UserID" json:"posts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var ExampleUser = User{
	Username:  "john_doe",
	Email:     "john.doe@example.com",
	Password:  "securepassword123", // Ensure this is hashed in a real application
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
