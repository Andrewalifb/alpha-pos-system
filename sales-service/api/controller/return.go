package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosReturnController interface {
	HandleCreatePosReturnRequest(c *gin.Context)
	HandleReadPosReturnRequest(c *gin.Context)
	HandleUpdatePosReturnRequest(c *gin.Context)
	HandleDeletePosReturnRequest(c *gin.Context)
	HandleReadAllPosReturnsRequest(c *gin.Context)
}

type posReturnController struct {
	service pb.PosReturnServiceClient
}

func NewPosReturnController(service pb.PosReturnServiceClient) PosReturnController {
	return &posReturnController{
		service: service,
	}
}
