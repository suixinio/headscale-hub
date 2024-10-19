package hs_router

import (
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/controller/hs_controller"
)

func InitNodeRoutes(r *gin.RouterGroup) gin.IRoutes {
	Group := r.Group("/node")
	controller := hs_controller.NewNodeController()
	Group.GET("/list", controller.List)
	Group.POST("/register", controller.Register)
	Group.POST("/rename", controller.Rename)
	Group.DELETE("/delete", controller.Delete)
	Group.POST("/move", controller.Move)
	return r
}
