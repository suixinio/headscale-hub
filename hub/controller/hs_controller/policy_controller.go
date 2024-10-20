package hs_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/suixinio/headscale-hub/common"
	"github.com/suixinio/headscale-hub/repository/hs_repository"
	"github.com/suixinio/headscale-hub/response"
	"github.com/suixinio/headscale-hub/vo"
)

type IPolicyController interface {
	GetAcl(c *gin.Context) // 获取acl
	SetAcl(c *gin.Context) // 设置acl
}

type PolicyController struct {
	PolicyRepository hs_repository.IPolicyRepository
}

func NewPolicyController() IPolicyController {
	hs_repository.NewPolicyRepository()
	PolicyController := PolicyController{PolicyRepository: hs_repository.NewPolicyRepository()}
	return PolicyController
}

// GetAcl 获取acl
func (pc PolicyController) GetAcl(c *gin.Context) {
	policy, err := pc.PolicyRepository.GetPolicy()
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{"data": policy}, "success")
}

// SetAcl 设置acl
func (pc PolicyController) SetAcl(c *gin.Context) {
	var req vo.SetACLRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	// 获取
	if err := pc.PolicyRepository.SetPolicy(req.Content); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "保存成功")
}
