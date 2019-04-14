package models

type MicroTaskBlackList struct {
	Id     int64 `gorm:"pk autoincr comment('id') BIGINT(20)"`
	UserId int64 `gorm:"not null comment('用户id') BIGINT(20)"`
}

func (c *MicroTaskBlackList) TableName() string {
	return "micro_task_black_list"
}
