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

type IRouteController interface {
	List(c *gin.Context)
	EnableStatus(c *gin.Context)
	Delete(c *gin.Context)
}

type RouteController struct {
	RouteRepository  hs_repository.IRouteRepository
	UserRepository   repository.IUserRepository
	HsUserRepository hs_repository.IUserRepository
}

func NewRouteController() IRouteController {
	RouteController := RouteController{
		RouteRepository:  hs_repository.NewRouteRepository(),
		UserRepository:   repository.NewUserRepository(),
		HsUserRepository: hs_repository.NewUserRepository(),
	}
	return RouteController
}

// List 获取所有路由
func (rc RouteController) List(c *gin.Context) {
	routes, err := rc.RouteRepository.ListRoutes()
	if err != nil {
		response.Fail(c, nil, "Failed to get Routes")
		return
	}
	response.Success(c, gin.H{"routes": routes}, "success")
}

// Status 启用路由
func (rc RouteController) EnableStatus(c *gin.Context) {
	var req vo.StatusRoute
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
	if err := rc.RouteRepository.SwitchStatus(req.RouteId, req.Enable); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{}, "success")
}

// Delete 删除路由
func (rc RouteController) Delete(c *gin.Context) {
	var req vo.DeleteRoute

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
	err := rc.RouteRepository.DeleteRoute(&req.DeleteRouteRequest)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{}, "success")
}
