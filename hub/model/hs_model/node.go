package hs_model

type NodeID uint64

// Node is a Headscale client.
type Node struct {
	ID   NodeID `gorm:"primary_key"`
	User User   `gorm:"constraint:OnDelete:CASCADE;"`
}

type (
	Nodes []*Node
)

func (hs *Node) TableName() string {
	return "nodes"
}
