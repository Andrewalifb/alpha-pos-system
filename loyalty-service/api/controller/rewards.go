package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosRewardController interface {
	HandleCreatePosRewardRequest(c *gin.Context)
	HandleReadPosRewardRequest(c *gin.Context)
	HandleUpdatePosRewardRequest(c *gin.Context)
	HandleDeletePosRewardRequest(c *gin.Context)
	HandleReadAllPosRewardsRequest(c *gin.Context)
}

type posRewardController struct {
	service pb.PosRewardServiceClient
}

func NewPosRewardController(service pb.PosRewardServiceClient) PosRewardController {
	return &posRewardController{
		service: service,
	}
}
