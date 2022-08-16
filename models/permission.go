package models

import (
	"time"
)

// Permission holds specific application settings linked to an Account.
type Permission struct {
	ID        int       `json:"-"`
	AccountID int       `json:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Theme string `json:"theme,omitempty"`
}
