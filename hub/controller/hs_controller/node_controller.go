package hs_controller

import "github.com/gin-gonic/gin"

type INodeController interface {
	List(c *gin.Context)
	Register(c *gin.Context)
	Rename(c *gin.Context)
	Delete(c *gin.Context)
	Move(c *gin.Context)
}

type NodeController struct {
}

func NewNodeController() INodeController {
	nodeController := NodeController{}
	return nodeController
}

// List 获取节点
func (nc NodeController) List(c *gin.Context) {

}

// Register 注册节点
func (nc NodeController) Register(c *gin.Context) {

}

// Rename 重命令节点
func (nc NodeController) Rename(c *gin.Context) {

}

// Delete 删除节点
func (nc NodeController) Delete(c *gin.Context) {

}

func (nc NodeController) Move(c *gin.Context) {

}
