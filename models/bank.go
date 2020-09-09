package models

import "github.com/google/uuid"

// Bank represent the bank model in the database
type Bank struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid,default:gen_random_uuid(),primary_key"`
	Name   string    `json:"name"`
	Active bool      `json:"-"`
}
