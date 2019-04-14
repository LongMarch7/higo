package models

type MicroCasbinRule struct {
	PType string `gorm:"VARCHAR(255)"`
	V0    string `gorm:"VARCHAR(255)"`
	V1    string `gorm:"VARCHAR(255)"`
	V2    string `gorm:"VARCHAR(255)"`
	V3    string `gorm:"VARCHAR(255)"`
	V4    string `gorm:"VARCHAR(255)"`
	V5    string `gorm:"VARCHAR(255)"`
}

func (c *MicroCasbinRule) TableName() string {
	return "micro_casbin_rule"
}
