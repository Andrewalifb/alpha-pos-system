package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosOnlinePaymentController interface {
	HandleCreatePosOnlinePaymentRequest(c *gin.Context)
	HandleReadPosOnlinePaymentRequest(c *gin.Context)
	HandleUpdatePosOnlinePaymentRequest(c *gin.Context)
	HandleDeletePosOnlinePaymentRequest(c *gin.Context)
	HandleReadAllPosOnlinePaymentsRequest(c *gin.Context)
}

type posOnlinePaymentController struct {
	service pb.PosOnlinePaymentServiceClient
}

func NewPosOnlinePaymentController(service pb.PosOnlinePaymentServiceClient) PosOnlinePaymentController {
	return &posOnlinePaymentController{
		service: service,
	}
}
