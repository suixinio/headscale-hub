package hs_controller

import "github.com/gin-gonic/gin"

type IRouteController interface {
	List(c *gin.Context)
	Enable(c *gin.Context)
	Disable(c *gin.Context)
	Delete(c *gin.Context)
}

type RouteController struct {
}

func NewRouteController() IRouteController {
	RouteController := RouteController{}
	return RouteController
}

// List 获取所有路由
func (rc RouteController) List(c *gin.Context) {

}

// Enable 启用路由
func (rc RouteController) Enable(c *gin.Context) {

}

// Disable 禁用路由
func (rc RouteController) Disable(c *gin.Context) {

}

// Delete 删除路由
func (rc RouteController) Delete(c *gin.Context) {

}
