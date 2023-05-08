package repository

import (
	"context"
	"fmt"
	entity "layout-item/internal/vip/model"
	"layout-item/pkg/db"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

type VipDbRepoI interface {
	QueryVipGradeExpConfList(ctx context.Context, param *entity.VipGradeExpConfListQuery, page, limit uint32) (int64, []entity.VipGradeExpConf, error)
	GetTestFmt()
}

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	db.NewDBConnMap,
	wire.Bind(new(db.DbConnI), new(*db.DbConnMap)),
)

var _ VipDbRepoI = (*VipDbStore)(nil)

type VipDbStore struct {
	Conn db.DbConnI
}

func NewVipDbStore(conn db.DbConnI) *VipDbStore {
	return &VipDbStore{
		Conn: conn,
	}
}

// 获取vip等级经验配置列表
func (v *VipDbStore) QueryVipGradeExpConfList(ctx context.Context, param *entity.VipGradeExpConfListQuery, page, limit uint32) (int64, []entity.VipGradeExpConf, error) {
	result := []entity.VipGradeExpConf{}
	dbConn := v.Conn.GetCtxDBConn(ctx)

	var total int64
	if err := dbConn.Model(&entity.VipGradeExpConf{}).Scopes(param.Where()).Count(&total).Error; err != nil {
		return 0, nil, errors.Wrap(err, "获取vip等级经验配置列表总数")
	}

	if err := dbConn.Scopes(param.Where(), db.Pagination(page, limit)).Find(&result).Error; err != nil {
		return 0, nil, errors.Wrap(err, "获取vip等级经验配置列表")
	}

	return total, result, nil
}

func (v *VipDbStore) GetTestFmt() {
	fmt.Println("db repo")
}
