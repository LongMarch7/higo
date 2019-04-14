package models

type MicroPortalTag struct {
	Id          int64  `gorm:"pk autoincr comment('分类id') BIGINT(20)"`
	Status      int    `gorm:"not null default 1 comment('状态,1:发布,0:不发布') TINYINT(3)"`
	Recommended int    `gorm:"not null default 0 comment('是否推荐;1:推荐;0:不推荐') TINYINT(3)"`
	PostCount   int64  `gorm:"not null default 0 comment('标签文章数') BIGINT(20)"`
	Name        string `gorm:"not null default '' comment('标签名称') VARCHAR(20)"`
}

func (c *MicroPortalTag) TableName() string {
	return "micro_portal_tag"
}
