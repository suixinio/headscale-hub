package hs_repository

type IPreAuthKeyRepository interface {
	//GetPreAuthKey() (*hsmodel.PreAuthKey, error)
	//SetPreAuthKey(c *gin.Context, content string) error
}

type PreAuthKeyRepository struct {
}

// HsPreAuthKeyRepository构造函数
func NewPreAuthKeyRepository() IPreAuthKeyRepository {
	return &PreAuthKeyRepository{}
}
