package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosPromotionController interface {
	HandleCreatePosPromotionRequest(c *gin.Context)
	HandleReadPosPromotionRequest(c *gin.Context)
	HandleUpdatePosPromotionRequest(c *gin.Context)
	HandleDeletePosPromotionRequest(c *gin.Context)
	HandleReadAllPosPromotionsRequest(c *gin.Context)
}

type posPromotionController struct {
	service pb.PosPromotionServiceClient
}

func NewPosPromotionController(service pb.PosPromotionServiceClient) PosPromotionController {
	return &posPromotionController{
		service: service,
	}
}
