package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosRewardRedemptionController interface {
	HandleCreatePosRewardRedemptionRequest(c *gin.Context)
	HandleReadPosRewardRedemptionRequest(c *gin.Context)
	HandleUpdatePosRewardRedemptionRequest(c *gin.Context)
	HandleDeletePosRewardRedemptionRequest(c *gin.Context)
	HandleReadAllPosRewardRedemptionsRequest(c *gin.Context)
}

type posRewardRedemptionController struct {
	service pb.PosRewardRedemptionServiceClient
}

func NewPosRewardRedemptionController(service pb.PosRewardRedemptionServiceClient) PosRewardRedemptionController {
	return &posRewardRedemptionController{
		service: service,
	}
}
