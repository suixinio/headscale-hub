package hs_repository

import (
	"context"
	pb "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/suixinio/headscale-hub/common"
	"github.com/suixinio/headscale-hub/model/hs_model"
)

type INodeRepository interface {
	ListNodes(user string) ([]*pb.Node, error)
	RegisterNode(user string, key string) (*pb.Node, error)
	DeleteNode(nodeId uint64) error
	CheckNodeRole(user *hs_model.User, nodeId uint64) bool
	RenameNode(nodeId uint64, name string) error
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

func (nr *NodeRepository) DeleteNode(nodeId uint64) error {
	_, err := common.HeadscaleGRPC.DeleteNode(context.Background(), &pb.DeleteNodeRequest{NodeId: nodeId})
	if err != nil {
		return err
	}
	return nil
}

func (nr *NodeRepository) CheckNodeRole(user *hs_model.User, nodeId uint64) bool {
	if err := common.DB.
		Order("id DESC").
		Limit(1).
		First(&hs_model.Node{User: *user, ID: hs_model.NodeID(nodeId)}).Error; err != nil {
		return false
	}
	return true
}

func (nr *NodeRepository) RenameNode(nodeId uint64, name string) error {
	_, err := common.HeadscaleGRPC.RenameNode(context.Background(), &pb.RenameNodeRequest{
		NodeId:  nodeId,
		NewName: name,
	})
	return err
}
