package hs_repository

import (
	"errors"
	"github.com/suixinio/headscale-hub/common"
	hsmodel "github.com/suixinio/headscale-hub/model/hs_model"
)

type IUserRepository interface {
	GetUserByName(name string) (*hsmodel.User, error)
}

type UserRepository struct {
}

// NewUserRepository 构造函数
func NewUserRepository() IUserRepository {
	return UserRepository{}
}

func (hs UserRepository) GetUserByName(name string) (*hsmodel.User, error) {
	var firstUser hsmodel.User
	err := common.DB.
		Where("name = ?", name).
		First(&firstUser).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &firstUser, nil
}
