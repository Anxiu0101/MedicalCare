package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	UID  uint `json:"uid"`
	User User `json:"user" gorm:"foreignKey:UID"`

	Tags []Tag `json:"tag" gorm:"foreignKey:name"`

	Title   string `json:"title" gorm:"size:255"`
	Desc    string `json:"desc" gorm:"size:255"`
	Content string `json:"content" gorm:"type:text"`
	State   int    `json:"state" gorm:"default:1"`
}
