package models

import "time"

// Contact is the contacts model
type Contact struct {
	ID        int
	First     string
	Last      string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
