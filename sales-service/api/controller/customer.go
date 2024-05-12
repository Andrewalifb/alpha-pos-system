package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosCustomerController interface {
	HandleCreatePosCustomerRequest(c *gin.Context)
	HandleReadPosCustomerRequest(c *gin.Context)
	HandleUpdatePosCustomerRequest(c *gin.Context)
	HandleDeletePosCustomerRequest(c *gin.Context)
	HandleReadAllPosCustomersRequest(c *gin.Context)
}

type posCustomerController struct {
	service pb.PosCustomerServiceClient
}

func NewPosCustomerController(service pb.PosCustomerServiceClient) PosCustomerController {
	return &posCustomerController{
		service: service,
	}
}
