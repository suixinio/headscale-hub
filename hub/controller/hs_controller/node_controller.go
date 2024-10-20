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

type INodeController interface {
	List(c *gin.Context)
	Register(c *gin.Context)
	Rename(c *gin.Context)
	Delete(c *gin.Context)
	Move(c *gin.Context)
}

type NodeController struct {
	NodeRepository   hs_repository.INodeRepository
	UserRepository   repository.IUserRepository
	HsUserRepository hs_repository.IUserRepository
}

func NewNodeController() INodeController {
	nodeController := NodeController{
		NodeRepository:   hs_repository.NewNodeRepository(),
		UserRepository:   repository.NewUserRepository(),
		HsUserRepository: hs_repository.NewUserRepository(),
	}
	return nodeController
}

// List 获取节点
func (nc NodeController) List(c *gin.Context) {
	role, user, err := nc.UserRepository.GetCurrentUserMinRoleSort(c)
	if err != nil {
		response.Fail(c, nil, "Failed to get Node")
		return
	}

	if role == 1 {
		user.Username = ""
	}
	nodes, err := nc.NodeRepository.ListNodes(user.Username)
	if err != nil {
		response.Fail(c, nil, "Failed to get Nodes")
		return
	}
	response.Success(c, gin.H{"nodes": nodes}, "success")
}

// Register 注册节点
func (nc NodeController) Register(c *gin.Context) {
	var req vo.RegisterNode
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
	user, err := nc.UserRepository.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, "Register to Node")
		return
	}
	node, err := nc.NodeRepository.RegisterNode(user.Username, req.Key)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{"node": node}, "success")
}

// Rename 重命名节点
func (nc NodeController) Rename(c *gin.Context) {

}

// Delete 删除节点
func (nc NodeController) Delete(c *gin.Context) {
	var req vo.DeleteNode
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
	role, user, err := nc.UserRepository.GetCurrentUserMinRoleSort(c)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	hsUser, err := nc.HsUserRepository.GetUserByName(user.Username)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 判断当前用户是否有权限
	if role != 1 {
		if !nc.NodeRepository.CheckNodeRole(hsUser, req.NodeId) {
			response.Fail(c, nil, "Delete to Node")
		}
	}

	err = nc.NodeRepository.DeleteNode(req.NodeId)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{}, "success")
}

// Move
func (nc NodeController) Move(c *gin.Context) {

}
