package models

import (
	"time"
)

// ArtWork holds specific application settings linked to an Account.
type ArtWork struct {
	ID        int       `json:"-"`
	AccountID int       `json:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Theme string `json:"theme,omitempty"`
}
