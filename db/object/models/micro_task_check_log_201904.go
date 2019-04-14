package models

type MicroTaskCheckLog201904 struct {
	Id             int64  `gorm:"pk autoincr comment('id') BIGINT(20)"`
	UserId         int64  `gorm:"not null default 1 comment('用户id') BIGINT(20)"`
	TaskId         int64  `gorm:"not null default 1 comment('任务日志id') BIGINT(20)"`
	ReasonDescribe string `gorm:"not null default '' comment('原因描述') VARCHAR(200)"`
}

func (c *MicroTaskCheckLog201904) TableName() string {
	return "micro_task_check_log_201904"
}
