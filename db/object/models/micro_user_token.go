package models

type MicroUserToken struct {
	Id         int64  `gorm:"pk autoincr BIGINT(20)"`
	UserId     int64  `gorm:"not null default 0 comment('用户id') BIGINT(20)"`
	ExpireTime int    `gorm:"not null default 0 comment('过期时间') INT(10)"`
	CreateTime int    `gorm:"not null default 0 comment('创建时间') INT(10)"`
	Token      string `gorm:"not null default '' comment('token') VARCHAR(64)"`
	DeviceType string `gorm:"not null default '' comment('设备类型;mobile,android,iphone,ipad,web,pc,mac,wxapp') VARCHAR(10)"`
}

func (c *MicroUserToken) TableName() string {
	return "micro_user_token"
}
