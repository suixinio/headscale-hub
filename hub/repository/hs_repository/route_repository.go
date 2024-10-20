package hs_repository

import (
	"context"
	pb "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/suixinio/headscale-hub/common"
)

type IRouteRepository interface {
	ListRoutes() ([]*pb.Route, error)
	DeleteRoute(req *pb.DeleteRouteRequest) error
	SwitchStatus(routeId uint64, enable bool) error
}

type RouteRepository struct {
}

// NewRouteRepository 构造函数
func NewRouteRepository() IRouteRepository {
	return &RouteRepository{}
}

func (rr *RouteRepository) ListRoutes() ([]*pb.Route, error) {
	resp, err := common.HeadscaleGRPC.GetRoutes(context.Background(), &pb.GetRoutesRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Routes, nil
}
func (rr *RouteRepository) DeleteRoute(req *pb.DeleteRouteRequest) error {
	_, err := common.HeadscaleGRPC.DeleteRoute(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func (rr *RouteRepository) SwitchStatus(routeId uint64, enable bool) error {
	var err error
	if enable {
		_, err = common.HeadscaleGRPC.EnableRoute(context.Background(), &pb.EnableRouteRequest{RouteId: routeId})
	} else {
		_, err = common.HeadscaleGRPC.DisableRoute(context.Background(), &pb.DisableRouteRequest{RouteId: routeId})
	}
	return err
}
