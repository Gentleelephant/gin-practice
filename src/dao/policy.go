package dao

import (
	"gin-practice/src/entity"
	"gorm.io/gorm"
)

type CasbinRuleManager struct {
	db *gorm.DB
}

func (p *CasbinRuleManager) CreatePolicy(policy *entity.CasbinRule) error {
	return p.db.Create(policy).Error
}
