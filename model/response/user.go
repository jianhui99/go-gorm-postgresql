package response

import (
	models "jd99/model"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID        `json:"id,omitempty"`
	Name      string           `json:"name,omitempty"`
	Email     string           `json:"email,omitempty"`
	Role      string           `json:"role,omitempty"`
	Photo     string           `json:"photo,omitempty"`
	Provider  string           `json:"provider"`
	Addresses []models.Address `json:"addresses"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
