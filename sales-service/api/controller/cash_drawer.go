package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosCashDrawerController interface {
	HandleCreatePosCashDrawerRequest(c *gin.Context)
	HandleReadPosCashDrawerRequest(c *gin.Context)
	HandleUpdatePosCashDrawerRequest(c *gin.Context)
	HandleDeletePosCashDrawerRequest(c *gin.Context)
	HandleReadAllPosCashDrawersRequest(c *gin.Context)
}

type posCashDrawerController struct {
	service pb.PosCashDrawerServiceClient
}

func NewPosCashDrawerController(service pb.PosCashDrawerServiceClient) PosCashDrawerController {
	return &posCashDrawerController{
		service: service,
	}
}
