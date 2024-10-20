package hs_model

import "gorm.io/gorm"

// User is the way Headscale implements the concept of users in Tailscale
//
// At the end of the day, users in Tailscale are some kind of 'bubbles' or users
// that contain our machines.
type User struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func (hs *User) TableName() string {
	return "users"
}
