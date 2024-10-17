package hs_controller

type IPreAuthKeyController interface {
}

type PreAuthKeyController struct {
}

func NewPreAuthKeyController() IPreAuthKeyController {
	//apiRepository := repository.NewApiRepository()
	PreAuthKeyController := PreAuthKeyController{}
	return PreAuthKeyController
}
