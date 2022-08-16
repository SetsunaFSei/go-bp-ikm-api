package models

import (
	"time"
)

// ChannelType holds specific application settings linked to an Account.
type ChannelType struct {
	ID        int       `json:"-"`
	AccountID int       `json:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Theme string `json:"theme,omitempty"`
}
