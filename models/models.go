package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Description string `json:"description" gorm:"text; not null; default:null`
	Completed  bool `json:"completed" gorm:"not null; default:false`
}