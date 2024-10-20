package vo

import pb "github.com/juanfont/headscale/gen/go/headscale/v1"

type SetACLRequest struct {
	Content string `json:"content" validate:"required"`
}

type RegisterNode struct {
	Key string `json:"key" validate:"required"`
}

type CreatePreAuthKey struct {
	pb.CreatePreAuthKeyRequest
}
