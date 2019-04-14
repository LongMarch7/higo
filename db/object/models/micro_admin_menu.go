package models

type MicroAdminMenu struct {
	Id         int     `gorm:"not null pk INT(10)"`
	ParentId   int     `gorm:"not null default 0 comment('父菜单id') index INT(10)"`
	Type       int     `gorm:"not null default 1 comment('菜单类型;1:有界面可访问菜单,2:无界面可访问菜单,0:只作为菜单') TINYINT(3)"`
	Status     int     `gorm:"not null default 0 comment('状态;1:显示,0:不显示') index TINYINT(3)"`
	ListOrder  float32 `gorm:"not null default 10000 comment('排序') FLOAT"`
	App        string  `gorm:"not null default '' comment('应用名') VARCHAR(40)"`
	Controller string  `gorm:"not null default '' comment('控制器名') index VARCHAR(30)"`
	Action     string  `gorm:"not null default '' comment('操作名称') VARCHAR(30)"`
	Param      string  `gorm:"not null default '' comment('额外参数') VARCHAR(50)"`
	Name       string  `gorm:"not null default '' comment('菜单名称') VARCHAR(30)"`
	Icon       string  `gorm:"not null default '' comment('菜单图标') VARCHAR(20)"`
	Remark     string  `gorm:"not null default '' comment('备注') VARCHAR(255)"`
}

func (c *MicroAdminMenu) TableName() string {
	return "micro_admin_menu"
}
