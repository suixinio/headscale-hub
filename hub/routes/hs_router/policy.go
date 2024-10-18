package hs_router

import (
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/controller/hs_controller"
)

func InitPolicyRoutes(r *gin.RouterGroup) gin.IRoutes {
	Group := r.Group("/policy")
	policyController := hs_controller.NewPolicyController()
	Group.GET("/acl", policyController.GetAcl)
	Group.POST("/acl", policyController.SetAcl)
	return r
}
