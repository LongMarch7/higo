package models

type MicroPortalCategoryPost struct {
	Id         int64   `gorm:"pk autoincr BIGINT(20)"`
	PostId     int64   `gorm:"not null default 0 comment('文章id') BIGINT(20)"`
	CategoryId int64   `gorm:"not null default 0 comment('分类id') index BIGINT(20)"`
	ListOrder  float32 `gorm:"not null default 10000 comment('排序') FLOAT"`
	Status     int     `gorm:"not null default 1 comment('状态,1:发布;0:不发布') TINYINT(3)"`
}

func (c *MicroPortalCategoryPost) TableName() string {
	return "micro_portal_category_post"
}
