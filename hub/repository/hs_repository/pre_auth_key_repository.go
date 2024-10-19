package hs_repository

import (
	"context"
	pb "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/suixinio/headscale-hub/common"
)

type IPreAuthKeyRepository interface {
	ListPreAuthKeys(user string) ([]*pb.PreAuthKey, error)
	//RegisterNode(user string, key string) (*pb.Node, error)
}

type PreAuthKeyRepository struct {
}

// NewPreAuthKeyRepository 构造函数
func NewPreAuthKeyRepository() IPreAuthKeyRepository {
	return &PreAuthKeyRepository{}
}

func (pr *PreAuthKeyRepository) ListPreAuthKeys(user string) ([]*pb.PreAuthKey, error) {
	keys, err := common.HeadscaleGRPC.ListPreAuthKeys(context.Background(), &pb.ListPreAuthKeysRequest{User: user})
	if err != nil {
		return nil, err
	}
	return keys.PreAuthKeys, nil
}
