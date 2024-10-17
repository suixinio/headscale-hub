package headscale_router

import (
	"go-web-mini/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	Group := r.Group("/headscale")
	Group.Use(authMiddleware.MiddlewareFunc())
	Group.Use(middleware.CasbinMiddleware())

	// 节点
	InitNodeRoutes(Group)
	// 路由
	InitRouteRoutes(Group)
	// acl
	InitPolicyRoutes(Group)
	// 密钥
	InitPreAuthKeyRoutes(Group)
	return r
}
