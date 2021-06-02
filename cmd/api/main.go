package main

import (
	"context"
	"os"

	"github.com/chanhlab/go-service-template/config"
	"github.com/chanhlab/go-service-template/internal/grpc"
	"github.com/chanhlab/go-utils/logger"
)

// Server ...
type Server struct {
	Config *config.Config
}

// RunServer ... runs gRPC server and HTTP gateway
func (srv *Server) RunServer() error {
	ctx := context.Background()

	// initialize logger
	logger.Init(srv.Config.Logger.LogLevel, srv.Config.Logger.LogTimeFormat)

	return grpc.RunGrpcServer(ctx, srv.Config)
}

func main() {
	config.NewConfig()
	server := &Server{
		Config: config.AppConfig,
	}

	if err := server.RunServer(); err != nil {
		logger.Log.Sugar().Error(err)
		os.Exit(1)
	}
}
