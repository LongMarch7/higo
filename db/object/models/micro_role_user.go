package models

type MicroRoleUser struct {
	Id     int64 `gorm:"pk autoincr BIGINT(20)"`
	RoleId int   `gorm:"not null default 0 comment('角色 id') index INT(10)"`
	UserId int64 `gorm:"not null default 0 comment('用户id') index BIGINT(20)"`
}

func (c *MicroRoleUser) TableName() string {
	return "micro_role_user"
}
