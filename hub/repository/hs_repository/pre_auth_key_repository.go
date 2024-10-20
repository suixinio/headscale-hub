package hs_repository

import (
	"context"
	pb "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/suixinio/headscale-hub/common"
)

type IPreAuthKeyRepository interface {
	ListPreAuthKeys(user string) ([]*pb.PreAuthKey, error)
	CreatePreAuthKey(request *pb.CreatePreAuthKeyRequest) (*pb.PreAuthKey, error)
	ExpirePreAuthKey(request *pb.ExpirePreAuthKeyRequest) (*pb.ExpirePreAuthKeyResponse, error)
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

func (pr *PreAuthKeyRepository) CreatePreAuthKey(request *pb.CreatePreAuthKeyRequest) (*pb.PreAuthKey, error) {
	resp, err := common.HeadscaleGRPC.CreatePreAuthKey(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp.PreAuthKey, nil
}

func (pr *PreAuthKeyRepository) ExpirePreAuthKey(request *pb.ExpirePreAuthKeyRequest) (*pb.ExpirePreAuthKeyResponse, error) {
	resp, err := common.HeadscaleGRPC.ExpirePreAuthKey(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
