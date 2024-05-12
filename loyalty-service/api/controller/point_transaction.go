package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosPointsTransactionController interface {
	HandleCreatePosPointsTransactionRequest(c *gin.Context)
	HandleReadPosPointsTransactionRequest(c *gin.Context)
	HandleUpdatePosPointsTransactionRequest(c *gin.Context)
	HandleDeletePosPointsTransactionRequest(c *gin.Context)
	HandleReadAllPosPointsTransactionsRequest(c *gin.Context)
}

type posPointsTransactionController struct {
	service pb.PosPointsTransactionServiceClient
}

func NewPosPointsTransactionController(service pb.PosPointsTransactionServiceClient) PosPointsTransactionController {
	return &posPointsTransactionController{
		service: service,
	}
}
