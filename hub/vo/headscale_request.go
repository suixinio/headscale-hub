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

type ExpirePreAuthKey struct {
	pb.ExpirePreAuthKeyRequest
}

type DeleteNode struct {
	pb.DeleteNodeRequest
}

type DeleteRoute struct {
	pb.DeleteRouteRequest
}

type StatusRoute struct {
	RouteId uint64 `json:"route_id" validate:"required"`
	Enable  bool   `json:"enable"`
}
