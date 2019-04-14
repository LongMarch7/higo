package models

type MicroUserLoginAttempt struct {
	Id            int64  `gorm:"pk autoincr BIGINT(20)"`
	LoginAttempts int    `gorm:"not null default 0 comment('尝试次数') INT(10)"`
	AttemptTime   int    `gorm:"not null default 0 comment('尝试登录时间') INT(10)"`
	LockedTime    int    `gorm:"not null default 0 comment('锁定时间') INT(10)"`
	Ip            string `gorm:"not null default '' comment('用户 ip') VARCHAR(15)"`
	Account       string `gorm:"not null default '' comment('用户账号,手机号,邮箱或用户名') VARCHAR(100)"`
}

func (c *MicroUserLoginAttempt) TableName() string {
	return "micro_user_login_attempt"
}
