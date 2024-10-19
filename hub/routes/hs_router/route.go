package hs_router

import (
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/controller/hs_controller"
)

func InitRouteRoutes(r *gin.RouterGroup) gin.IRoutes {
	Group := r.Group("/route")
	controller := hs_controller.NewRouteController()
	Group.GET("/list", controller.List)
	Group.GET("/enable", controller.Enable)
	Group.GET("/disable", controller.Disable)
	Group.GET("/delete", controller.Delete)
	return r
}
