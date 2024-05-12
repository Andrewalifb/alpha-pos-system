package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosSaleController interface {
	HandleCreatePosSaleRequest(c *gin.Context)
	HandleReadPosSaleRequest(c *gin.Context)
	HandleUpdatePosSaleRequest(c *gin.Context)
	HandleDeletePosSaleRequest(c *gin.Context)
	HandleReadAllPosSalesRequest(c *gin.Context)
}

type posSaleController struct {
	service pb.PosSaleServiceClient
}

func NewPosSaleController(service pb.PosSaleServiceClient) PosSaleController {
	return &posSaleController{
		service: service,
	}
}
