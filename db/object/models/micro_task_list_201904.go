package models

type MicroTaskList201904 struct {
	Id           int64   `gorm:"pk autoincr comment('id') BIGINT(20)"`
	UserId       int64   `gorm:"not null default 1 comment('发布者id') BIGINT(20)"`
	ClassId      int64   `gorm:"not null default 1 comment('分类id') BIGINT(20)"`
	Count        int     `gorm:"not null default 1 comment('任务数量') INT(15)"`
	ConsumeCount int     `gorm:"not null default 0 comment('已接单任务数量') INT(15)"`
	CheckCount   int     `gorm:"not null default 0 comment('提交审核任务数量') INT(15)"`
	FinishCount  int     `gorm:"not null default 0 comment('完成任务数量') INT(15)"`
	Balance      string  `gorm:"not null default 0.00 comment('任务单价') DECIMAL(10,2)"`
	CreateTime   int     `gorm:"not null default 0 comment('发布日期') INT(11)"`
	EndTime      int     `gorm:"not null default 0 comment('结束日期') INT(11)"`
	ListOrder    float32 `gorm:"not null default 10000 comment('排序') FLOAT"`
	RegionRule   int64   `gorm:"not null default 0 comment('地域id限制') BIGINT(10)"`
	ClassRule    int     `gorm:"not null default 0 comment('类别限制;0: 不限制，1: 同类别限制') TINYINT(2)"`
	Status       int     `gorm:"not null default 0 comment('发布状态;0:未发布,1: 审核中,2:已发布') TINYINT(3)"`
	TaskTitle    string  `gorm:"not null default '' comment('任务标题') VARCHAR(100)"`
	TaskDescribe string  `gorm:"not null default '' comment('任务简述') VARCHAR(200)"`
	Thumbnail    string  `gorm:"not null default '' comment('缩略图') VARCHAR(100)"`
	TaskContent  string  `gorm:"comment('任务内容') TEXT"`
}

func (c *MicroTaskList201904) TableName() string {
	return "micro_task_list_201904"
}
