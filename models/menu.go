package models

import (
	"time"
)

// Menu holds specific application settings linked to an Account.
type Menu struct {
	ID        int       `json:"-"`
	AccountID int       `json:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Theme string `json:"theme,omitempty"`
}
