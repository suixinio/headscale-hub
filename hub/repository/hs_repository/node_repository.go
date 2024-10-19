package hs_repository

import (
	"context"
	pb "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/suixinio/headscale-hub/common"
)

type INodeRepository interface {
	ListNodes(user string) ([]*pb.Node, error)
	RegisterNode(user string, key string) (*pb.Node, error)
}

type NodeRepository struct {
}

// NewNodeRepository 构造函数
func NewNodeRepository() INodeRepository {
	return &NodeRepository{}
}

func (nr *NodeRepository) ListNodes(user string) ([]*pb.Node, error) {
	nodes, err := common.HeadscaleGRPC.ListNodes(context.Background(), &pb.ListNodesRequest{User: user})
	if err != nil {
		return nil, err
	}
	return nodes.Nodes, nil
}

func (nr *NodeRepository) RegisterNode(user string, key string) (*pb.Node, error) {
	node, err := common.HeadscaleGRPC.RegisterNode(context.Background(), &pb.RegisterNodeRequest{User: user, Key: key})
	if err != nil {
		return nil, err
	}
	return node.Node, err
}
