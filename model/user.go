package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex;not null" json:"username"`
	Email    string    `gorm:"uniqueIndex;not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	Names    string    `json:"names"`
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
