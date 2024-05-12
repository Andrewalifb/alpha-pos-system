package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosInventoryHistoryController interface {
	HandleCreatePosInventoryHistoryRequest(c *gin.Context)
	HandleReadPosInventoryHistoryRequest(c *gin.Context)
	HandleUpdatePosInventoryHistoryRequest(c *gin.Context)
	HandleDeletePosInventoryHistoryRequest(c *gin.Context)
	HandleReadAllPosInventoryHistoriesRequest(c *gin.Context)
}

type posInventoryHistoryController struct {
	service pb.PosInventoryHistoryServiceClient
}

func NewPosInventoryHistoryController(service pb.PosInventoryHistoryServiceClient) PosInventoryHistoryController {
	return &posInventoryHistoryController{
		service: service,
	}
}
