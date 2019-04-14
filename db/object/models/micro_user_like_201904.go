package models

type MicroUserLike201904 struct {
	Id          int64  `gorm:"pk autoincr BIGINT(20)"`
	UserId      int64  `gorm:"not null default 0 comment('用户 id') index BIGINT(20)"`
	ObjectId    int    `gorm:"not null default 0 comment('内容原来的主键id') INT(10)"`
	CreateTime  int    `gorm:"not null default 0 comment('创建时间') INT(10)"`
	TabName     string `gorm:"not null default '' comment('内容以前所在表,不带前缀') VARCHAR(64)"`
	Url         string `gorm:"not null default '' comment('内容的原文地址，不带域名') VARCHAR(255)"`
	Title       string `gorm:"not null default '' comment('内容的标题') VARCHAR(100)"`
	Thumbnail   string `gorm:"not null default '' comment('缩略图') VARCHAR(100)"`
	Description string `gorm:"comment('内容的描述') TEXT"`
}

func (c *MicroUserLike201904) TableName() string {
	return "micro_user_like_201904"
}
