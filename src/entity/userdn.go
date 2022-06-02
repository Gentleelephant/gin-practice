package entity

import "gorm.io/gorm"

type UserDn struct {
	gorm.Model
	Username string `json:"username"`
	Dn       string `json:"dn"`
}
