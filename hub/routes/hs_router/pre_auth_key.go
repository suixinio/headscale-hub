package hs_router

import (
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/controller/hs_controller"
)

func InitPreAuthKeyRoutes(r *gin.RouterGroup) gin.IRoutes {
	Group := r.Group("/pre_auth_key")
	controller := hs_controller.NewPreAuthKeyController()
	Group.GET("/list", controller.List)
	Group.POST("/create", controller.Create)
	Group.DELETE("/expire", controller.Expire)
	return r
}
