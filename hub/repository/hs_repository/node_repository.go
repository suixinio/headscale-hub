package hs_repository

type INodeRepository interface {
	//GetPreAuthKey() (*hsmodel.PreAuthKey, error)
	//SetPreAuthKey(c *gin.Context, content string) error
}

type NodeRepository struct {
}

// HsPreAuthKeyRepository构造函数
func NewNodeRepository() INodeRepository {
	return &NodeRepository{}
}
