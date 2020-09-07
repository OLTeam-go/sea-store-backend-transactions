package models

import "github.com/google/uuid"

// Bank represent the bank model in the database
type Bank struct {
	ID     uuid.UUID `json:"id" pg:"type:uuid,default:gen_random_uuid(),pk"`
	Name   string    `json:"name"`
	Active string    `json:"active"`
}
