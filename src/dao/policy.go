package dao

import (
	"gin-practice/src/entity"
	"gorm.io/gorm"
)

type CasbinRuleManager struct {
	DB *gorm.DB
}

func (p *CasbinRuleManager) CreatePolicy(policy *entity.CasbinRule) error {
	return p.DB.Create(policy).Error
}
