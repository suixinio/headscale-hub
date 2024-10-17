package hs_controller

type IPolicyController interface {
}

type PolicyController struct {
}

func NewPolicyController() IPolicyController {
	//apiRepository := repository.NewApiRepository()
	PolicyController := PolicyController{}
	return PolicyController
}
