package hs_repository

import (
	ccontext "context"
	pb "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/suixinio/headscale-hub/common"
	"golang.org/x/net/context"
)

type IPolicyRepository interface {
	GetPolicy() (string, error)
	SetPolicy(content string) error
}

type PolicyRepository struct {
}

// NewPolicyRepository 构造函数
func NewPolicyRepository() IPolicyRepository {
	return &PolicyRepository{}
}

// GetPolicy returns the latest policy in the database.
func (hs *PolicyRepository) GetPolicy() (string, error) {
	res, err := common.HeadscaleGRPC.GetPolicy(ccontext.Background(), &pb.GetPolicyRequest{})
	if err != nil {
		return "", err
	}
	return res.Policy, nil
}

// SetPolicy Updates the ACL Policy.
func (hs *PolicyRepository) SetPolicy(content string) error {
	_, err := common.HeadscaleGRPC.SetPolicy(context.Background(), &pb.SetPolicyRequest{Policy: content})
	if err != nil {
		return err
	}
	return nil
}
