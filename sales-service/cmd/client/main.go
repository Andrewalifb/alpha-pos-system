package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Andrewalifb/alpha-pos-system/sales-service/api/controller"
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/sales-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50054", "the address to connect to")
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
	cashDrawerClient := pb.NewPosCashDrawerServiceClient(conn)
	customerClient := pb.NewPosCustomerServiceClient(conn)
	invoiceClient := pb.NewPosInvoiceServiceClient(conn)
	onlinePaymentClient := pb.NewPosOnlinePaymentServiceClient(conn)
	paymentMethodClient := pb.NewPosPaymentMethodServiceClient(conn)
	returnClient := pb.NewPosReturnServiceClient(conn)
	saleClient := pb.NewPosSaleServiceClient(conn)

	// Initialize the controllers with the gRPC clients
	cashDrawerCtrl := controller.NewPosCashDrawerController(cashDrawerClient)
	customerCtrl := controller.NewPosCustomerController(customerClient)
	invoiceCtrl := controller.NewPosInvoiceController(invoiceClient)
	onlinePaymentCtrl := controller.NewPosOnlinePaymentController(onlinePaymentClient)
	paymentMethodCtrl := controller.NewPosPaymentMethodController(paymentMethodClient)
	returnCtrl := controller.NewPosReturnController(returnClient)
	saleCtrl := controller.NewPosSaleController(saleClient)

	// Create a new router
	r := gin.Default()

	// Define your routes
	routes.PosCashDrawerRoutes(r, cashDrawerCtrl)
	routes.PosCustomerRoutes(r, customerCtrl)
	routes.PosInvoiceRoutes(r, invoiceCtrl)
	routes.PosOnlinePaymentRoutes(r, onlinePaymentCtrl)
	routes.PosPaymentMethodRoutes(r, paymentMethodCtrl)
	routes.PosReturnRoutes(r, returnCtrl)
	routes.PosSaleRoutes(r, saleCtrl)

	// Start the server
	r.Run(":8083")
}
