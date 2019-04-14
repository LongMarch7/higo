package models

type MicroComment struct {
	Id           int64  `gorm:"pk autoincr BIGINT(20)"`
	ParentId     int64  `gorm:"not null default 0 comment('被回复的评论id') index BIGINT(20)"`
	UserId       int    `gorm:"not null default 0 comment('发表评论的用户id') INT(10)"`
	ToUserId     int    `gorm:"not null default 0 comment('被评论的用户id') INT(10)"`
	ObjectId     int    `gorm:"not null default 0 comment('评论内容 id') index index(table_id_status) INT(10)"`
	LikeCount    int    `gorm:"not null default 0 comment('点赞数') INT(10)"`
	DislikeCount int    `gorm:"not null default 0 comment('不喜欢数') INT(10)"`
	Floor        int    `gorm:"not null default 0 comment('楼层数') INT(10)"`
	CreateTime   int    `gorm:"not null default 0 comment('评论时间') index INT(10)"`
	DeleteTime   int    `gorm:"not null default 0 comment('删除时间') INT(10)"`
	Status       int    `gorm:"not null default 1 comment('状态,1:已审核,0:未审核') index index(table_id_status) TINYINT(3)"`
	Type         int    `gorm:"not null default 1 comment('评论类型；1实名评论') TINYINT(3)"`
	TabName      string `gorm:"not null default '' comment('评论内容所在表，不带表前缀') index(table_id_status) VARCHAR(64)"`
	FullName     string `gorm:"not null default '' comment('评论者昵称') VARCHAR(50)"`
	Email        string `gorm:"not null default '' comment('评论者邮箱') VARCHAR(255)"`
	Path         string `gorm:"not null default '' comment('层级关系') VARCHAR(255)"`
	Url          string `gorm:"comment('原文地址') TEXT"`
	Content      string `gorm:"comment('评论内容') TEXT"`
	More         string `gorm:"comment('扩展属性') TEXT"`
}

func (c *MicroComment) TableName() string {
	return "micro_comment"
}
