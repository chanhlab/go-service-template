// Package grpc ...
package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/chanhlab/go-service-template/config"
	"github.com/chanhlab/go-utils/grpc/middleware"
	"github.com/chanhlab/go-utils/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunGrpcServer ...
func RunGrpcServer(ctx context.Context, appConfig *config.Config) error {
	logger.Log.Sugar().Info(fmt.Sprintf("gRPC Port: %d", appConfig.Server.GRPCPort))
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.Server.GRPCPort))
	if err != nil {
		logger.Log.Sugar().Error(err)
		return err
	}

	// gRPC server statup options
	options := []grpc.ServerOption{}

	// add middleware
	options = middleware.AddLogging(logger.Log, options)

	// register server
	server := grpc.NewServer(options...)

	// Get DB connection
	// db := mysql.GetConnection(
	// 	appConfig.MySQL.Host,
	// 	appConfig.MySQL.Port,
	// 	appConfig.MySQL.DBName,
	// 	appConfig.MySQL.Username,
	// 	appConfig.MySQL.Password,
	// 	appConfig.MySQL.MaxIDLEConnection,
	// 	appConfig.MySQL.MaxOpenConnection,
	// )

	// if db == nil {
	// 	return errors.New("can not create mysql connection")
	// }

	// New Repository
	// userRepository := models.NewUserRepository(db)

	// New Service
	// userService := services.NewUserService(userRepository)

	// Register service
	// userv1.RegisterUserAPIServer(server, userService)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	reflection.Register(server)
	err = server.Serve(listen)
	if err != nil {
		logger.Log.Sugar().Error(err)
		return err
	}
	return server.Serve(listen)
}
