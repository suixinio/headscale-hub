package common

import (
	"context"
	"fmt"
	pb "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/suixinio/headscale-hub/config"
	"google.golang.org/grpc"
)

type headscaleGRPC struct {
	pb.HeadscaleServiceClient
	opts []grpc.DialOption
}

var HeadscaleGRPC *headscaleGRPC

func InitHeadscaleGrpc() {
	h := &headscaleGRPC{}
	h.opts = []grpc.DialOption{grpc.WithPerRPCCredentials(h), grpc.WithInsecure()}

	conn, err := grpc.NewClient(config.Conf.Headscale.Host, h.opts...)
	if err != nil {
		panic(fmt.Sprintf("did not connect: %v", err))
	}
	h.HeadscaleServiceClient = pb.NewHeadscaleServiceClient(conn)
	HeadscaleGRPC = h
}

// GetRequestMetadata gets the current request metadata, refreshing tokens
func (h *headscaleGRPC) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + config.Conf.Headscale.ApiKey,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
func (h *headscaleGRPC) RequireTransportSecurity() bool {
	return false
}
