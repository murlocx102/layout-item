package entity

import "gorm.io/gorm"

type VipGradeExpConfListQuery struct {
	ID int64 // 等级配置ID
}

// Where 条件
func (v VipGradeExpConfListQuery) Where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v.ID != 0 {
			db = db.Where("id = ?", v.ID)
		}

		return db
	}
}
