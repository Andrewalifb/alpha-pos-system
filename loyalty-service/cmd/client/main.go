package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/controller"
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create gRPC clients for each service
	customerLoyaltyClient := pb.NewPosCustomerLoyaltyServiceClient(conn)
	loyaltyLevelClient := pb.NewPosLoyaltyLevelServiceClient(conn)
	pointsTransactionClient := pb.NewPosPointsTransactionServiceClient(conn)
	rewardRedemptionClient := pb.NewPosRewardRedemptionServiceClient(conn)
	rewardClient := pb.NewPosRewardServiceClient(conn)

	// Initialize the controllers with the gRPC clients
	customerLoyaltyCtrl := controller.NewPosCustomerLoyaltyController(customerLoyaltyClient)
	loyaltyLevelCtrl := controller.NewPosLoyaltyLevelController(loyaltyLevelClient)
	pointsTransactionCtrl := controller.NewPosPointsTransactionController(pointsTransactionClient)
	rewardRedemptionCtrl := controller.NewPosRewardRedemptionController(rewardRedemptionClient)
	rewardCtrl := controller.NewPosRewardController(rewardClient)

	// Create a new router
	r := gin.Default()

	// Define your routes
	routes.PosCustomerLoyaltyRoutes(r, customerLoyaltyCtrl)
	routes.PosLoyaltyLevelRoutes(r, loyaltyLevelCtrl)
	routes.PosPointsTransactionRoutes(r, pointsTransactionCtrl)
	routes.PosRewardRedemptionRoutes(r, rewardRedemptionCtrl)
	routes.PosRewardRoutes(r, rewardCtrl)

	// Start the server
	r.Run(":8081")
}
