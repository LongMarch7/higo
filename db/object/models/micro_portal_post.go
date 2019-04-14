package models

type MicroPortalPost struct {
	Id                  int64  `gorm:"pk autoincr index(type_status_date) BIGINT(20)"`
	ParentId            int64  `gorm:"not null default 0 comment('父级id') index BIGINT(20)"`
	PostType            int    `gorm:"not null default 1 comment('类型,1:文章;2:页面') index(type_status_date) TINYINT(3)"`
	PostFormat          int    `gorm:"not null default 1 comment('内容格式;1:html;2:md') TINYINT(3)"`
	UserId              int64  `gorm:"not null default 0 comment('发表者用户id') index BIGINT(20)"`
	PostStatus          int    `gorm:"not null default 1 comment('状态;1:已发布;0:未发布;') index(type_status_date) TINYINT(3)"`
	CommentStatus       int    `gorm:"not null default 1 comment('评论状态;1:允许;0:不允许') TINYINT(3)"`
	IsTop               int    `gorm:"not null default 0 comment('是否置顶;1:置顶;0:不置顶') TINYINT(3)"`
	Recommended         int    `gorm:"not null default 0 comment('是否推荐;1:推荐;0:不推荐') TINYINT(3)"`
	PostHits            int64  `gorm:"not null default 0 comment('查看数') BIGINT(20)"`
	PostFavorites       int    `gorm:"not null default 0 comment('收藏数') INT(10)"`
	PostLike            int64  `gorm:"not null default 0 comment('点赞数') BIGINT(20)"`
	CommentCount        int64  `gorm:"not null default 0 comment('评论数') BIGINT(20)"`
	CreateTime          int    `gorm:"not null default 0 comment('创建时间') index index(type_status_date) INT(10)"`
	UpdateTime          int    `gorm:"not null default 0 comment('更新时间') INT(10)"`
	PublishedTime       int    `gorm:"not null default 0 comment('发布时间') INT(10)"`
	DeleteTime          int    `gorm:"not null default 0 comment('删除时间') INT(10)"`
	PostTitle           string `gorm:"not null default '' comment('post标题') VARCHAR(100)"`
	PostKeywords        string `gorm:"not null default '' comment('seo keywords') VARCHAR(150)"`
	PostExcerpt         string `gorm:"not null default '' comment('post摘要') VARCHAR(500)"`
	PostSource          string `gorm:"not null default '' comment('转载文章的来源') VARCHAR(150)"`
	Thumbnail           string `gorm:"not null default '' comment('缩略图') VARCHAR(100)"`
	PostContent         string `gorm:"comment('文章内容') TEXT"`
	PostContentFiltered string `gorm:"comment('处理过的文章内容') TEXT"`
	More                string `gorm:"comment('扩展属性,如缩略图;格式为json') TEXT"`
}

func (c *MicroPortalPost) TableName() string {
	return "micro_portal_post"
}
