package models

type MicroUserFavorite struct {
	Id          int64  `gorm:"pk autoincr BIGINT(20)"`
	UserId      int64  `gorm:"not null default 0 comment('用户 id') index BIGINT(20)"`
	Title       string `gorm:"not null default '' comment('收藏内容的标题') VARCHAR(100)"`
	Thumbnail   string `gorm:"not null default '' comment('缩略图') VARCHAR(100)"`
	Url         string `gorm:"comment('收藏内容的原文地址，JSON格式') VARCHAR(255)"`
	Description string `gorm:"comment('收藏内容的描述') TEXT"`
	TabName     string `gorm:"not null default '' comment('收藏实体以前所在表,不带前缀') VARCHAR(64)"`
	ObjectId    int    `gorm:"default 0 comment('收藏内容原来的主键id') INT(10)"`
	CreateTime  int    `gorm:"default 0 comment('收藏时间') INT(10)"`
}

func (c *MicroUserFavorite) TableName() string {
	return "micro_user_favorite"
}
