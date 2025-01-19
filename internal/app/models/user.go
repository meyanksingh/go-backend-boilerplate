package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName    string    `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName     string    `json:"last_name" gorm:"type:varchar(100);not null"`
	Email        string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password     string    `json:"password" gorm:"type:varchar(255);not null"`
	Token        string    `json:"token" gorm:"type:varchar(255)"`
	RefreshToken string    `json:"refresh_token" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
