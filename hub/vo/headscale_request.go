package vo

type SetACLRequest struct {
	Content string `json:"content" validate:"required"`
}
