package app

import (
	"context"
	"log/slog"

	"gprc_kafka_streamer/internal/broker"
	"gprc_kafka_streamer/proto/v1/service"
)

type GRPCServer struct {
	service.UnimplementedKafkaStreamerServer
	Broker broker.Broker
	Logger *slog.Logger
}

func NewServer(msgBroker broker.Broker, logger *slog.Logger) *GRPCServer {
	return &GRPCServer{
		Broker: msgBroker,
		Logger: logger,
	}
}

func (s *GRPCServer) CreateChannel(
	ctx context.Context,
	in *service.CreateChannelRequest,
) (*service.CreateChannelResponse, error) {
	s.Logger.Info("Method CreateChannel not implemented yet")
	//TODO: implement method
	return &service.CreateChannelResponse{}, nil
}
