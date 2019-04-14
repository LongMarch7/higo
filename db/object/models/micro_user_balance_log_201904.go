package models

type MicroUserBalanceLog201904 struct {
	Id          int64  `gorm:"pk autoincr BIGINT(20)"`
	UserId      int64  `gorm:"not null default 0 comment('用户 id') BIGINT(20)"`
	CreateTime  int    `gorm:"not null default 0 comment('创建时间') INT(10)"`
	Change      string `gorm:"not null default 0.00 comment('更改余额') DECIMAL(20,2)"`
	Balance     string `gorm:"not null default 0.00 comment('更改后余额') DECIMAL(20,2)"`
	Fee         string `gorm:"not null default 0.00 comment('手续费') DECIMAL(10,2)"`
	Description string `gorm:"not null default '' comment('描述') VARCHAR(255)"`
	Remark      string `gorm:"not null default '' comment('备注') VARCHAR(255)"`
}

func (c *MicroUserBalanceLog201904) TableName() string {
	return "micro_user_balance_log_201904"
}
