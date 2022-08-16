package models

import (
	"time"
)

// Quota holds specific application settings linked to an Account.
type Quota struct {
	ID        int       `json:"-"`
	AccountID int       `json:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Theme string `json:"theme,omitempty"`
}
