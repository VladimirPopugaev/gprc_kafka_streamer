package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gprc_kafka_streamer/internal/app"
	kf "gprc_kafka_streamer/internal/broker/kafka"
	"gprc_kafka_streamer/pkg/config"
	sl "gprc_kafka_streamer/pkg/logger"
	"gprc_kafka_streamer/proto/v1/service"
)

func main() {
	// create context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// read config
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config is not found. Error: %s", err)
	}

	// create logger
	logger, err := sl.SetupLogger(cfg.Env)
	if err != nil {
		log.Fatalf("Logger is not created. Error: %s", err)
	}

	// Create kafka client
	broker, err := kf.New(ctx, cfg.Broker.Address)
	if err != nil {
		log.Fatalf("Kafka client is not created. Error: %s", err)
	}

	// create service
	srv := app.NewServer(broker, logger)

	// start service
	s := grpc.NewServer()
	service.RegisterKafkaStreamerServer(s, srv)

	go func() {
		listen, err := net.Listen("tcp", cfg.Server.Address)
		if err != nil {
			log.Fatalf("TCP listen error: %s", err)
		}

		if err := s.Serve(listen); err != nil {
			log.Fatalf("Service Kafka Streamer start error: %s", err)
		}
	}()

	log.Printf("Service Kafka Streamer started. Address: %s", cfg.Server.Address)

	<-ctx.Done()
	log.Printf("Service Kafka Streamer stopped")

	// 3 second for stopping all process
	timeCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.GracefulStop()

	<-timeCtx.Done()
	log.Printf("Program correctly finished")
}
