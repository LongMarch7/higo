package models

type MicroRole struct {
	Id         int     `gorm:"not null pk autoincr INT(10)"`
	ParentId   int     `gorm:"not null default 0 comment('父角色ID') index INT(10)"`
	Status     int     `gorm:"not null default 0 comment('状态;0:禁用;1:正常') index TINYINT(3)"`
	CreateTime int     `gorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int     `gorm:"not null default 0 comment('更新时间') INT(10)"`
	ListOrder  float32 `gorm:"not null default 0 comment('排序') FLOAT"`
	Name       string  `gorm:"not null default '' comment('角色名称') VARCHAR(20)"`
	Remark     string  `gorm:"not null default '' comment('备注') VARCHAR(255)"`
}

func (c *MicroRole) TableName() string {
	return "micro_role"
}
