package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosLoyaltyLevelController interface {
	HandleCreatePosLoyaltyLevelRequest(c *gin.Context)
	HandleReadPosLoyaltyLevelRequest(c *gin.Context)
	HandleUpdatePosLoyaltyLevelRequest(c *gin.Context)
	HandleDeletePosLoyaltyLevelRequest(c *gin.Context)
	HandleReadAllPosLoyaltyLevelsRequest(c *gin.Context)
}

type posLoyaltyLevelController struct {
	service pb.PosLoyaltyLevelServiceClient
}

func NewPosLoyaltyLevelController(service pb.PosLoyaltyLevelServiceClient) PosLoyaltyLevelController {
	return &posLoyaltyLevelController{
		service: service,
	}
}
