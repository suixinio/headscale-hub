package hs_router

import (
	"github.com/gin-gonic/gin"
)

func InitNodeRoutes(r *gin.RouterGroup) gin.IRoutes {
	Group := r.Group("/node")
	Group.GET("/machine")
	Group.POST("/machine")
	Group.DELETE("/machine")
	Group.PUT("/machine")
	Group.PATCH("/machine")
	return r
}
