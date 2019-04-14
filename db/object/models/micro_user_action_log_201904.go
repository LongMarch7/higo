package models

type MicroUserActionLog201904 struct {
	Id            int64  `gorm:"pk autoincr BIGINT(20)"`
	UserId        int64  `gorm:"not null default 0 comment('用户id') index(user_object_action) index(user_object_action_ip) BIGINT(20)"`
	Count         int    `gorm:"not null default 0 comment('访问次数') INT(10)"`
	LastVisitTime int    `gorm:"not null default 0 comment('最后访问时间') INT(10)"`
	Object        string `gorm:"not null default '' comment('访问对象的id,格式:不带前缀的表名+id;如posts1表示xx_posts表里id为1的记录') index(user_object_action) index(user_object_action_ip) VARCHAR(100)"`
	Action        string `gorm:"not null default '' comment('操作名称;格式:应用名+控制器+操作名,也可自己定义格式只要不发生冲突且惟一;') index(user_object_action) index(user_object_action_ip) VARCHAR(50)"`
	Ip            string `gorm:"not null default '' comment('用户ip') index(user_object_action_ip) VARCHAR(15)"`
}

func (c *MicroUserActionLog201904) TableName() string {
	return "micro_user_action_log_201904"
}
