package entity

import "layout-item/pkg/db"

// vip等级经验配置
type VipGradeExpConf struct {
	db.GormModel
	Grade        int64  `json:"grade" gorm:"column:grade"`                   // vip等级
	GradeName    string `json:"grade_name" gorm:"column:grade_name"`         // 等级名称
	Icon         string `json:"icon" gorm:"column:icon"`                     // 等级图标
	Exp          int64  `json:"exp" gorm:"column:exp"`                       // 经验值
	DayDeductExp int64  `json:"day_deduct_exp" gorm:"column:day_deduct_exp"` // 每日预扣除基础值
	RebateRatio  int64  `json:"rebate_ratio" gorm:"column:rebate_ratio"`     // 返利比例(单位:万分比)
}

// TableName 确定表名
func (v *VipGradeExpConf) TableName() string {
	return "vip_grade_exp_conf"
}
