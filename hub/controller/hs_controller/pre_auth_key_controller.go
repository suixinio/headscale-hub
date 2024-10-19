package hs_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/repository"
	"github.com/suixinio/headscale-hub/repository/hs_repository"
	"github.com/suixinio/headscale-hub/response"
)

type IPreAuthKeyController interface {
	List(c *gin.Context)   // 获取密钥列表
	Create(c *gin.Context) // 创建密钥
	Expire(c *gin.Context) // 密钥过期
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
	if err != nil && err.Error() != "rpc error: code = Unknown desc = User not found" {
		fmt.Println(err)
		response.Fail(c, nil, "Failed to get Nodes")
		return
	}
	response.Success(c, gin.H{"keys": keys}, "success")
}

// Create 创建密钥
func (pc PreAuthKeyController) Create(c *gin.Context) {
}

// Expire 密钥过期
func (pc PreAuthKeyController) Expire(c *gin.Context) {
}
