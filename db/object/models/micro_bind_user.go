package models

type MicroBindUser struct {
	Id         int64  `gorm:"pk autoincr comment('id') BIGINT(20)"`
	UserId     string `gorm:"not null comment('绑定账号') VARCHAR(20)"`
	Start      int    `gorm:"not null default 2 comment('粉丝数') INT(11)"`
	RegionId   int    `gorm:"not null default 0 comment('用户地域') INT(10)"`
	ClassId    int    `gorm:"not null default 0 comment('用户分类') INT(10)"`
	BindStatus int    `gorm:"not null default 0 comment('绑定状态;0:未绑定,1: 审核中,2:已绑定') TINYINT(3)"`
	PassTime   int    `gorm:"not null default 0 comment('审核通过日期') INT(11)"`
}

func (c *MicroBindUser) TableName() string {
	return "micro_bind_user"
}
