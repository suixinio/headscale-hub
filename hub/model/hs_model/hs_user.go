package hs_model

import "gorm.io/gorm"

// User is the way Headscale implements the concept of users in Tailscale
//
// At the end of the day, users in Tailscale are some kind of 'bubbles' or users
// that contain our machines.
type HsUser struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func (hs *HsUser) TableName() string {
	return "users"
}
