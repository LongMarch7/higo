package models

type MicroUser struct {
	Id                int64  `gorm:"pk autoincr BIGINT(20)"`
	BindId            int64  `gorm:"not null comment('bind user id') BIGINT(20)"`
	UserType          int    `gorm:"not null default 2 comment('用户类型;1:admin;2:大v;3:广告商') TINYINT(3)"`
	Sex               int    `gorm:"not null default 0 comment('性别;0:保密,1:男,2:女') TINYINT(2)"`
	Birthday          int    `gorm:"not null default 0 comment('生日') INT(11)"`
	LastLoginTime     int    `gorm:"not null default 0 comment('最后登录时间') INT(11)"`
	Score             int    `gorm:"not null default 0 comment('用户积分') INT(11)"`
	Coin              int    `gorm:"not null default 0 comment('金币') INT(10)"`
	Balance           string `gorm:"not null default 0.00 comment('余额') DECIMAL(10,2)"`
	CreateTime        int    `gorm:"not null default 0 comment('注册时间') INT(11)"`
	UserStatus        int    `gorm:"not null default 1 comment('用户状态;0:禁用,1:正常,2:未验证') TINYINT(3)"`
	UserLogin         string `gorm:"not null default '' comment('用户名') index VARCHAR(60)"`
	UserPass          string `gorm:"not null default '' comment('登录密码;micro_password加密') VARCHAR(64)"`
	UserNickname      string `gorm:"not null default '' comment('用户昵称') index VARCHAR(50)"`
	UserEmail         string `gorm:"not null default '' comment('用户登录邮箱') VARCHAR(100)"`
	UserUrl           string `gorm:"not null default '' comment('用户个人网址') VARCHAR(100)"`
	Avatar            string `gorm:"not null default '' comment('用户头像') VARCHAR(255)"`
	Signature         string `gorm:"not null default '' comment('个性签名') VARCHAR(255)"`
	LastLoginIp       string `gorm:"not null default '' comment('最后登录ip') VARCHAR(15)"`
	UserActivationKey string `gorm:"not null default '' comment('激活码') VARCHAR(60)"`
	Mobile            string `gorm:"not null default '' comment('中国手机不带国家代码，国际手机号格式为：国家代码-手机号') VARCHAR(20)"`
	More              string `gorm:"comment('扩展属性') TEXT"`
}

func (c *MicroUser) TableName() string {
	return "micro_user"
}
