package hs_router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/middleware"
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
