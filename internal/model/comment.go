package model

type Comment struct {
	GormModel
	ArticleID uint    `gorm:"not null" json:"article_id"`
	Article   Article `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"article"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	Content  string `gorm:"type:text;not null" json:"content"`
	ParentID *uint  `json:"parent_id"` // 如果为 nil，就表示是顶级评论

}
