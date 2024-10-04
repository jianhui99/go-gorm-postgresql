package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID        uuid.UUID `gorm:"not null" json:"user_id,omitempty"`
	State         string
	City          string
	Postcode      int
	DetailAddress string
}
