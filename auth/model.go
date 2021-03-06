package auth

type MicroCasbinRule struct {
    PType string `xorm:"VARCHAR(255)"`
    V0    string `xorm:"VARCHAR(255)"`
    V1    string `xorm:"VARCHAR(255)"`
    V2    string `xorm:"VARCHAR(255)"`
    V3    string `xorm:"VARCHAR(255)"`
    V4    string `xorm:"VARCHAR(255)"`
    V5    string `xorm:"VARCHAR(255)"`
}

func (c *MicroCasbinRule) TableName() string {
    return "micro_casbin_rule"
}
