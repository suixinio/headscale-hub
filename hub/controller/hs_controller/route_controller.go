package hs_controller

type IRouteController interface {
}

type RouteController struct {
}

func NewRouteController() IRouteController {
	//apiRepository := repository.NewApiRepository()
	RouteController := RouteController{}
	return RouteController
}
