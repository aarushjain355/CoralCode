package models

import (
	"time"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"` // autoIncrement for integer IDs
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// User represents a user in the system.
type User struct {
	// Embedding BaseModel for ID, CreatedAt, UpdatedAt fields
	// This means User will implicitly have ID, CreatedAt, and UpdatedAt.
	BaseModel

	// Specific fields for the User model
	Username string `json:"username" gorm:"type:varchar(50);unique;not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	// Password is marked with json:"-" to prevent it from being marshaled into JSON responses
	// In a real application, you'd store a hashed password, not plain text.
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	IsActive bool   `json:"is_active" gorm:"default:true"` // Example: for account activation/deactivation
}
