package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosStoreController interface {
	HandleCreatePosStoreRequest(c *gin.Context)
	HandleReadPosStoreRequest(c *gin.Context)
	HandleUpdatePosStoreRequest(c *gin.Context)
	HandleDeletePosStoreRequest(c *gin.Context)
	HandleReadAllPosStoresRequest(c *gin.Context)
}

type posStoreController struct {
	service pb.PosStoreServiceClient
}

func NewPosStoreController(service pb.PosStoreServiceClient) PosStoreController {
	return &posStoreController{
		service: service,
	}
}
