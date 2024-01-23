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
	ID       uuid.UUID `gorm:"type:uuid;"`
}

type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}
