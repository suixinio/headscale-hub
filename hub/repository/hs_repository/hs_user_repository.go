package hsrepository

import (
	"errors"
	"github.com/suixinio/headscale-hub/common"
	hsmodel "github.com/suixinio/headscale-hub/model/hs_model"
	"time"

	"github.com/patrickmn/go-cache"
)

type IHsUserRepository interface {
	// Login(user *model.User) (*model.User, error)       // 登录
	// ChangePwd(username string, newPasswd string) error // 更新密码

	// CreateUser(user *model.User) error                              // 创建用户
	// GetUserById(id uint) (model.User, error)                        // 获取单个用户
	// GetUsers(req *vo.UserListRequest) ([]*model.User, int64, error) // 获取用户列表
	// UpdateUser(user *model.User) error                              // 更新用户
	// BatchDeleteUserByIds(ids []uint) error                          // 批量删除

	// GetCurrentUser(c *gin.Context) (model.User, error)                  // 获取当前登录用户信息
	// GetCurrentUserMinRoleSort(c *gin.Context) (uint, model.User, error) // 获取当前用户角色排序最小值（最高等级角色）以及当前用户信息
	// GetUserMinRoleSortsByIds(ids []uint) ([]int, error)                 // 根据用户ID获取用户角色排序最小值

	// SetUserInfoCache(username string, user model.User) // 设置用户信息缓存
	// UpdateUserInfoCacheByRoleId(roleId uint) error     // 根据角色ID更新拥有该角色的用户信息缓存
	// ClearUserInfoCache()                               // 清理所有用户信息缓存
}

type HsUserRepository struct {
}

// 当前用户信息缓存，避免频繁获取数据库
var userInfoCache = cache.New(24*time.Hour, 48*time.Hour)

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
