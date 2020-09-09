package models

import (
	"time"

	"github.com/google/uuid"
)

//Base generalize data in repository
type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid,default:gen_random_uuid(),primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
