package models

type MicroAsset struct {
	Id            int64  `gorm:"pk autoincr BIGINT(20)"`
	UserId        int64  `gorm:"not null default 0 comment('用户id') BIGINT(20)"`
	FileSize      int64  `gorm:"not null default 0 comment('文件大小,单位B') BIGINT(20)"`
	CreateTime    int    `gorm:"not null default 0 comment('上传时间') INT(10)"`
	Status        int    `gorm:"not null default 1 comment('状态;1:可用,0:不可用') TINYINT(3)"`
	DownloadTimes int    `gorm:"not null default 0 comment('下载次数') INT(10)"`
	FileKey       string `gorm:"not null default '' comment('文件惟一码') VARCHAR(64)"`
	Filename      string `gorm:"not null default '' comment('文件名') VARCHAR(100)"`
	FilePath      string `gorm:"not null default '' comment('文件路径,相对于upload目录,可以为url') VARCHAR(100)"`
	FileMd5       string `gorm:"not null default '' comment('文件md5值') VARCHAR(32)"`
	FileSha1      string `gorm:"not null default '' VARCHAR(40)"`
	Suffix        string `gorm:"not null default '' comment('文件后缀名,不包括点') VARCHAR(10)"`
	More          string `gorm:"comment('其它详细信息,JSON格式') TEXT"`
}

func (c *MicroAsset) TableName() string {
	return "micro_asset"
}
