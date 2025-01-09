package model

type Category struct {
	GormModel
	Name     string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Articles []Article `gorm:"foreignKey:CategoryID" json:"articles"`
}
