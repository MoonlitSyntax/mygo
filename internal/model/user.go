package model

type User struct {
	GormModel
	Email    string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Username string    `gorm:"type:varchar(50);not null" json:"username"`
	Password string    `gorm:"type:varchar(255);not null" json:"password"`
	Role     string    `gorm:"type:varchar(50);default:'guest'" json:"role"`
	Articles []Article `gorm:"foreignKey:CategoryID" json:"articles"`
}
