package hs_router

import (
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/controller/hs_controller"
)

func InitPolicyRoutes(r *gin.RouterGroup) gin.IRoutes {
	Group := r.Group("/policy")
	controller := hs_controller.NewPolicyController()
	Group.GET("/get", controller.GetAcl)
	Group.POST("/set", controller.SetAcl)
	return r
}
