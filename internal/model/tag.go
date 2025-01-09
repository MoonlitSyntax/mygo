package model

type Tag struct {
	GormModel
	Name     string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Articles []Article `gorm:"many2many:article_tag;"`
}
