package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosCustomerLoyaltyController interface {
	HandleCreatePosCustomerLoyaltyRequest(c *gin.Context)
	HandleReadPosCustomerLoyaltyRequest(c *gin.Context)
	HandleUpdatePosCustomerLoyaltyRequest(c *gin.Context)
	HandleDeletePosCustomerLoyaltyRequest(c *gin.Context)
	HandleReadAllPosCustomerLoyaltiesRequest(c *gin.Context)
}

type posCustomerLoyaltyController struct {
	service pb.PosCustomerLoyaltyServiceClient
}

func NewPosCustomerLoyaltyController(service pb.PosCustomerLoyaltyServiceClient) PosCustomerLoyaltyController {
	return &posCustomerLoyaltyController{
		service: service,
	}
}
