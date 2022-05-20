package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model

	Name  string `json:"name" gorm:"size:64;uniqueIndex"`
	State int    `json:"state" gorm:"default:1"`
}
