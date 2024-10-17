package headscale_router

import (
	"github.com/gin-gonic/gin"
)

func InitPolicyRoutes(r *gin.RouterGroup) gin.IRoutes {
	Group := r.Group("/policy")
	// headscaleGroup.GET("/acl", aclc.GetAccessControl)
	// headscaleGroup.POST("/acl", aclc.SetAccessControl)
	Group.GET("/acl")
	Group.POST("/acl")
	return r
}
