package service

import (
	"layout-item/internal/vip/repository"
)

type VipRepoI interface {
	GetGradeList()
}

var _ VipRepoI = (*VipService)(nil)

type VipService struct {
	dbRepo repository.VipDbRepoI
}

func NewVipService(repo repository.VipDbRepoI) *VipService {
	return &VipService{
		dbRepo: repo,
	}
}

func (v *VipService) GetGradeList() {
	v.dbRepo.GetTestFmt()
	return
}
