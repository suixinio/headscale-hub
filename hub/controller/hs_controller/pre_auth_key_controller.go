package hs_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/repository/hs_repository"
)

type IPreAuthKeyController interface {
	List(c *gin.Context)   // 获取密钥列表
	Create(c *gin.Context) // 创建密钥
	Expire(c *gin.Context) // 密钥过期
}

type PreAuthKeyController struct {
	PreAuthKeyRepository hs_repository.IPreAuthKeyRepository
}

func NewPreAuthKeyController() IPreAuthKeyController {
	PreAuthKeyController := PreAuthKeyController{PreAuthKeyRepository: hs_repository.NewPreAuthKeyRepository()}
	return PreAuthKeyController
}

// List 获取密钥列表
func (pc PreAuthKeyController) List(c *gin.Context) {
}

// Create 创建密钥
func (pc PreAuthKeyController) Create(c *gin.Context) {
}

// Expire 密钥过期
func (pc PreAuthKeyController) Expire(c *gin.Context) {
}
