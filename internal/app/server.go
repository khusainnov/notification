package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/khusainnov/notification/internal/app/notificationservice"
	"github.com/khusainnov/notification/internal/config"
	"github.com/khusainnov/notification/internal/email"
	"github.com/khusainnov/notification/internal/rabbitmq"
	napi "github.com/khusainnov/notification/pkg/notificationapi/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, cfg *config.Config) error {
	s := grpc.NewServer()

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		return fmt.Errorf("error due listen addr, %w", err)
	}

	emailClient := email.NewClient(cfg)

	consumer, err := rabbitmq.NewConsumer(cfg, emailClient)
	if err != nil {
		return fmt.Errorf("cannot create consumer, %w", err)
	}

	notificationClient := notificationservice.NewClient(emailClient)

	napi.RegisterNotificationServiceServer(s, notificationClient)
	reflection.Register(s)

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	go gracefulShutDown(cfg, s, cancel)
	go consumer.Run(ctx)

	cfg.L.Info("starting listening service", zap.Any("PORT", cfg.GRPCAddr))
	if err = s.Serve(lis); err != nil {
		return fmt.Errorf("error due server grpc server, %w", err)
	}

	return nil
}

func gracefulShutDown(cfg *config.Config, s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	c := <-ch
	cfg.L.Info("Called graceful shutdown", zap.Any("SIGNAL", c))

	s.GracefulStop()
	cancel()
}
