package models

type MicroPortalTagPost struct {
	Id     int64 `gorm:"pk autoincr BIGINT(20)"`
	TagId  int64 `gorm:"not null default 0 comment('标签 id') BIGINT(20)"`
	PostId int64 `gorm:"not null default 0 comment('文章 id') index BIGINT(20)"`
	Status int   `gorm:"not null default 1 comment('状态,1:发布;0:不发布') TINYINT(3)"`
}

func (c *MicroPortalTagPost) TableName() string {
	return "micro_portal_tag_post"
}
