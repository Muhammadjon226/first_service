package main

import (
	"net"

	"github.com/Muhammadjon226/first_service/config"
	pbFirst "github.com/Muhammadjon226/first_service/genproto/first_service"
	"github.com/Muhammadjon226/first_service/pkg/db"
	"github.com/Muhammadjon226/first_service/pkg/logger"
	"github.com/Muhammadjon226/first_service/service"
	"github.com/Muhammadjon226/first_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "first-service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to firstgres error", logger.Error(err))
	}
	pgStorage := storage.NewStoragePg(connDB)

	firstService := service.NewFirstService(pgStorage, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pbFirst.RegisterFirstServiceServer(s, firstService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
