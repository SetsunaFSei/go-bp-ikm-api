package models

import (
	"time"
)

// Footer holds specific application settings linked to an Account.
type Footer struct {
	ID        int       `json:"-"`
	TITLE     int       `json:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Theme string `json:"theme,omitempty"`
}
