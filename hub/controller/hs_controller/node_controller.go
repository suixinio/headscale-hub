package hs_controller

type INodeController interface {
}

type NodeController struct {
}

func NewNodeController() INodeController {
	//apiRepository := repository.NewApiRepository()
	nodeController := NodeController{}
	return nodeController
}
