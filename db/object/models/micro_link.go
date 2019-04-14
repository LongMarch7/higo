package models

type MicroLink struct {
	Id          int64   `gorm:"pk autoincr BIGINT(20)"`
	Status      int     `gorm:"not null default 1 comment('状态;1:显示;0:不显示') index TINYINT(3)"`
	Rating      int     `gorm:"not null default 0 comment('友情链接评级') INT(11)"`
	ListOrder   float32 `gorm:"not null default 10000 comment('排序') FLOAT"`
	Description string  `gorm:"not null default '' comment('友情链接描述') VARCHAR(255)"`
	Url         string  `gorm:"not null default '' comment('友情链接地址') VARCHAR(255)"`
	Name        string  `gorm:"not null default '' comment('友情链接名称') VARCHAR(30)"`
	Image       string  `gorm:"not null default '' comment('友情链接图标') VARCHAR(100)"`
	Target      string  `gorm:"not null default '' comment('友情链接打开方式') VARCHAR(10)"`
	Rel         string  `gorm:"not null default '' comment('链接与网站的关系') VARCHAR(50)"`
}

func (c *MicroLink) TableName() string {
	return "micro_link"
}
