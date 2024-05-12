package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosPaymentMethodController interface {
	HandleCreatePosPaymentMethodRequest(c *gin.Context)
	HandleReadPosPaymentMethodRequest(c *gin.Context)
	HandleUpdatePosPaymentMethodRequest(c *gin.Context)
	HandleDeletePosPaymentMethodRequest(c *gin.Context)
	HandleReadAllPosPaymentMethodsRequest(c *gin.Context)
}

type posPaymentMethodController struct {
	service pb.PosPaymentMethodServiceClient
}

func NewPosPaymentMethodController(service pb.PosPaymentMethodServiceClient) PosPaymentMethodController {
	return &posPaymentMethodController{
		service: service,
	}
}
