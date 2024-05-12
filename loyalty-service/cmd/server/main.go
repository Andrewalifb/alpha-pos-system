package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/config"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/pkg/repository"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/pkg/service"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	// Initialize the database
	dbConfig := config.NewConfig()

	// Initialize the repositories
	customerLoyaltyRepo := repository.NewPosCustomerLoyaltyRepository(dbConfig.SQLDB, dbConfig.RedisDB)
	loyaltyLevelRepo := repository.NewPosLoyaltyLevelRepository(dbConfig.SQLDB, dbConfig.RedisDB)
	pointsTransactionRepo := repository.NewPosPointsTransactionRepository(dbConfig.SQLDB, dbConfig.RedisDB)
	rewardRedemptionRepo := repository.NewPosRewardRedemptionRepository(dbConfig.SQLDB, dbConfig.RedisDB)
	rewardRepo := repository.NewPosRewardRepository(dbConfig.SQLDB, dbConfig.RedisDB)

	// Initialize the services
	customerLoyaltySvc := service.NewPosCustomerLoyaltyService(customerLoyaltyRepo)
	loyaltyLevelSvc := service.NewPosLoyaltyLevelService(loyaltyLevelRepo)
	pointsTransactionSvc := service.NewPosPointsTransactionService(pointsTransactionRepo)
	rewardRedemptionSvc := service.NewPosRewardRedemptionService(rewardRedemptionRepo)
	rewardSvc := service.NewPosRewardService(rewardRepo)

	// Create a gRPC server
	s := grpc.NewServer()

	// Register the services with the gRPC server
	pb.RegisterPosCustomerLoyaltyServiceServer(s, customerLoyaltySvc)
	pb.RegisterPosLoyaltyLevelServiceServer(s, loyaltyLevelSvc)
	pb.RegisterPosPointsTransactionServiceServer(s, pointsTransactionSvc)
	pb.RegisterPosRewardRedemptionServiceServer(s, rewardRedemptionSvc)
	pb.RegisterPosRewardServiceServer(s, rewardSvc)

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
