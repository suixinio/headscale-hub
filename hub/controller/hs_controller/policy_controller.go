package hs_controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/suixinio/headscale-hub/repository/hs_repository"
	"github.com/suixinio/headscale-hub/response"
)

type IPolicyController interface {
	GetAcl(c *gin.Context) // 获取acl
	SetAcl(c *gin.Context) // 设置acl
}

type PolicyController struct {
	PolicyRepository hs_repository.IHsPolicyRepository
}

func NewPolicyController() IPolicyController {
	hs_repository.NewPolicyRepository()
	PolicyController := PolicyController{PolicyRepository: hs_repository.NewPolicyRepository()}
	return PolicyController
}

// 获取acl
func (pc PolicyController) GetAcl(c *gin.Context) {
	policyM, err := pc.PolicyRepository.GetPolicy()
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	data, err := json.Marshal(policyM.Data)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{"data": string(data)}, "success")
}

// 设置acl
func (pc PolicyController) SetAcl(c *gin.Context) {

}
