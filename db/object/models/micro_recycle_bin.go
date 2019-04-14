package models

type MicroRecycleBin struct {
	Id         int64  `gorm:"pk autoincr BIGINT(20)"`
	ObjectId   int    `gorm:"default 0 comment('删除内容 id') INT(11)"`
	CreateTime int    `gorm:"default 0 comment('创建时间') INT(10)"`
	TabName    string `gorm:"default '' comment('删除内容所在表名') VARCHAR(60)"`
	Name       string `gorm:"default '' comment('删除内容名称') VARCHAR(255)"`
	UserId     int64  `gorm:"not null default 0 comment('用户id') BIGINT(20)"`
}

func (c *MicroRecycleBin) TableName() string {
	return "micro_recycle_bin"
}
