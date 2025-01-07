package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
}
