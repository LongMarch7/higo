package models

type MicroUserScoreLog201904 struct {
	Id         int64  `gorm:"pk autoincr BIGINT(20)"`
	UserId     int64  `gorm:"not null default 0 comment('用户 id') BIGINT(20)"`
	CreateTime int    `gorm:"not null default 0 comment('创建时间') INT(11)"`
	Action     string `gorm:"not null default '' comment('用户操作名称') VARCHAR(50)"`
	Score      int    `gorm:"not null default 0 comment('更改积分，可以为负') INT(11)"`
	Coin       int    `gorm:"not null default 0 comment('更改金币，可以为负') INT(11)"`
}

func (c *MicroUserScoreLog201904) TableName() string {
	return "micro_user_score_log_201904"
}
