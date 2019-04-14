package models

type MicroAuthRule struct {
	Id     int    `gorm:"not null pk autoincr comment('规则id,自增主键') INT(10)"`
	Name   string `gorm:"not null default '' comment('规则唯一英文标识,全小写') VARCHAR(100)"`
	Method string `gorm:"not null default '' comment('规则方法get、post、put...') VARCHAR(100)"`
	Title  string `gorm:"not null default '' comment('规则描述') VARCHAR(20)"`
}

func (c *MicroAuthRule) TableName() string {
	return "micro_auth_rule"
}
