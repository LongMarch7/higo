package models

type MicroTaskCategory struct {
	Id             int64   `gorm:"pk autoincr comment('分类id') BIGINT(20)"`
	ParentId       int64   `gorm:"not null default 0 comment('分类父id') BIGINT(20)"`
	Count          int64   `gorm:"not null default 0 comment('分类任务数') BIGINT(20)"`
	Status         int     `gorm:"not null default 1 comment('状态,1:发布,0:不发布') TINYINT(3)"`
	DeleteTime     int     `gorm:"not null default 0 comment('删除时间') INT(10)"`
	ListOrder      float32 `gorm:"not null default 10000 comment('排序') FLOAT"`
	Name           string  `gorm:"not null default '' comment('分类名称') VARCHAR(200)"`
	Description    string  `gorm:"not null default '' comment('分类描述') VARCHAR(200)"`
	SeoTitle       string  `gorm:"not null default '' VARCHAR(100)"`
	SeoKeywords    string  `gorm:"not null default '' VARCHAR(200)"`
	SeoDescription string  `gorm:"not null default '' VARCHAR(200)"`
	OneTpl         string  `gorm:"not null default '' comment('分类模板') VARCHAR(50)"`
	More           string  `gorm:"comment('扩展属性') TEXT"`
}

func (c *MicroTaskCategory) TableName() string {
	return "micro_task_category"
}
