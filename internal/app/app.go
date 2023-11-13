package app

import (
	"context"
	"gprc_kafka_streamer/proto/v1/service"
)

type GRPCServer struct {
	service.UnimplementedKafkaStreamerServer
}

func (s *GRPCServer) CreateChannel(
	ctx context.Context,
	in *service.CreateChannelRequest,
) (*service.CreateChannelResponse, error) {
	return &service.CreateChannelResponse{}, nil
}
