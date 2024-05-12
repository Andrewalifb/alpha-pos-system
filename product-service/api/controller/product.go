package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosProductController interface {
	HandleCreatePosProductRequest(c *gin.Context)
	HandleReadPosProductRequest(c *gin.Context)
	HandleUpdatePosProductRequest(c *gin.Context)
	HandleDeletePosProductRequest(c *gin.Context)
	HandleReadAllPosProductsRequest(c *gin.Context)
}

type posProductController struct {
	service pb.PosProductServiceClient
}

func NewPosProductController(service pb.PosProductServiceClient) PosProductController {
	return &posProductController{
		service: service,
	}
}
