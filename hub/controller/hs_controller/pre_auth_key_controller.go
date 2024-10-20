package hs_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/suixinio/headscale-hub/common"
	"github.com/suixinio/headscale-hub/repository"
	"github.com/suixinio/headscale-hub/repository/hs_repository"
	"github.com/suixinio/headscale-hub/response"
	"github.com/suixinio/headscale-hub/vo"
)

type IPreAuthKeyController interface {
	List(c *gin.Context)   // 获取密钥列表
	Create(c *gin.Context) // 创建密钥
	Expire(c *gin.Context) // 密钥过期D
}

type PreAuthKeyController struct {
	PreAuthKeyRepository hs_repository.IPreAuthKeyRepository
	NodeRepository       hs_repository.INodeRepository
	UserRepository       repository.IUserRepository
}

func NewPreAuthKeyController() IPreAuthKeyController {
	PreAuthKeyController := PreAuthKeyController{
		PreAuthKeyRepository: hs_repository.NewPreAuthKeyRepository(),
		NodeRepository:       hs_repository.NewNodeRepository(),
		UserRepository:       repository.NewUserRepository(),
	}
	return PreAuthKeyController
}

// List 获取密钥列表
func (pc PreAuthKeyController) List(c *gin.Context) {
	user, err := pc.UserRepository.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, "Failed to get Node")
		return
	}
	keys, err := pc.PreAuthKeyRepository.ListPreAuthKeys(user.Username)
	if err != nil && err.Error() != "rpc error: code = Unknown desc = user not found" {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{"keys": keys}, "success")
}

// Create 创建密钥
func (pc PreAuthKeyController) Create(c *gin.Context) {
	var req vo.CreatePreAuthKey
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	user, err := pc.UserRepository.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	req.User = user.Username
	req.AclTags = []string{}
	key, err := pc.PreAuthKeyRepository.CreatePreAuthKey(&req.CreatePreAuthKeyRequest)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{"key": key}, "success")
}

// Expire 密钥过期
func (pc PreAuthKeyController) Expire(c *gin.Context) {
	var req vo.ExpirePreAuthKey
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	user, err := pc.UserRepository.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	req.User = user.Username
	_, err = pc.PreAuthKeyRepository.ExpirePreAuthKey(&req.ExpirePreAuthKeyRequest)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{}, "success")
}
