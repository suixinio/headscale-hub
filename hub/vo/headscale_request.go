package vo

type SetACLRequest struct {
	Content string `json:"content" validate:"required"`
}

type RegisterNode struct {
	Key string `json:"content" validate:"required"`
}
