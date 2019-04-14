package models

type MicroAuthAccess struct {
	Id       int64  `gorm:"pk autoincr BIGINT(20)"`
	RoleId   int    `gorm:"not null comment('角色') index INT(10)"`
	RuleName string `gorm:"not null default '' comment('规则唯一英文标识,全小写') index VARCHAR(100)"`
	Type     string `gorm:"not null default '' comment('权限规则分类,请加应用前缀,如admin_') VARCHAR(30)"`
}

func (c *MicroAuthAccess) TableName() string {
	return "micro_auth_access"
}
