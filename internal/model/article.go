package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Description string `gorm:"type:text" json:"description"`
	Top         bool   `gorm:"type:tinyint;default:0" json:"top"`
	BelongTo    string `gorm:"type:varchar(50);default:'public'" json:"belong_to"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"category"`

	Tags []Tag `gorm:"many2many:article_tag;" json:"tags"`
}
