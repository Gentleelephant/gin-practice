package entity

type CasbinRule struct {
	ID int `gorm:"primary_key autoincrement"`

	PType string `gorm:"column:ptype;type:varchar(100)"`

	V0 string `gorm:"column:v0;type:varchar(100)"`

	V1 string `gorm:"column:v1;type:varchar(100)"`

	V2 string `gorm:"column:v2;type:varchar(100)"`

	V3 string `gorm:"column:v3;type:varchar(100)"`

	V4 string `gorm:"column:v4;type:varchar(100)"`

	V5 string `gorm:"column:v5;type:varchar(100)"`

	V6 string `gorm:"column:v6;type:varchar(25)"`

	V7 string `gorm:"column:v7;type:varchar(25)"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
