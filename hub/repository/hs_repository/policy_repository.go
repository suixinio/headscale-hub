package hs_repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	pb "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/suixinio/headscale-hub/common"
	hsmodel "github.com/suixinio/headscale-hub/model/hs_model"

	"gorm.io/gorm"
)

type IPolicyRepository interface {
	GetPolicy() (*hsmodel.Policy, error)
	SetPolicy(c *gin.Context, content string) error
}

type PolicyRepository struct {
}

// HsPolicyRepository构造函数
func NewPolicyRepository() IPolicyRepository {
	return &PolicyRepository{}
}

// GetPolicy returns the latest policy in the database.
func (hs *PolicyRepository) GetPolicy() (*hsmodel.Policy, error) {
	var p hsmodel.Policy

	// Query:
	// SELECT * FROM policies ORDER BY id DESC LIMIT 1;
	if err := common.DB.
		Order("id DESC").
		Limit(1).
		First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, hsmodel.ErrPolicyNotFound
		}

		return nil, err
	}

	return &p, nil
}

// GetPolicy returns the latest policy in the database.
func (hs *PolicyRepository) SetPolicy(c *gin.Context, content string) error {
	_, err := common.HeadscaleGRPC.SetPolicy(c, &pb.SetPolicyRequest{Policy: content})
	if err != nil {
		return err
	}
	return nil
}
