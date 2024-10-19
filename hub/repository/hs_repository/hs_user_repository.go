package hs_repository

import (
	"errors"
	"github.com/suixinio/headscale-hub/common"
	hsmodel "github.com/suixinio/headscale-hub/model/hs_model"
)

type IHsUserRepository interface {
	GetUserByName(name string) (*hsmodel.HsUser, error)
}

type HsUserRepository struct {
}

// UserRepository构造函数
func NewHsUserRepository() HsUserRepository {
	return HsUserRepository{}
}

func (hs HsUserRepository) GetUserByName(name string) (*hsmodel.HsUser, error) {
	var firstUser hsmodel.HsUser
	err := common.DB.
		Where("name = ?", name).
		First(&firstUser).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &firstUser, nil
}
