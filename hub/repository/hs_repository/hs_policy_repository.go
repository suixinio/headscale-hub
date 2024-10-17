package hsrepository

import (
	"errors"
	"go-web-mini/common"
	hsmodel "go-web-mini/model/hs_model"

	"gorm.io/gorm"
)

type IHsPolicyRepository interface {
}

type HsPolicyRepository struct {
}

// HsPolicyRepository构造函数
func NewPolicyRepository() HsPolicyRepository {
	return HsPolicyRepository{}
}

// GetPolicy returns the latest policy in the database.
func (hs *HsPolicyRepository) GetPolicy() (*hsmodel.Policy, error) {
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
