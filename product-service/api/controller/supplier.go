package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosSupplierController interface {
	HandleCreatePosSupplierRequest(c *gin.Context)
	HandleReadPosSupplierRequest(c *gin.Context)
	HandleUpdatePosSupplierRequest(c *gin.Context)
	HandleDeletePosSupplierRequest(c *gin.Context)
	HandleReadAllPosSuppliersRequest(c *gin.Context)
}

type posSupplierController struct {
	service pb.PosSupplierServiceClient
}

func NewPosSupplierController(service pb.PosSupplierServiceClient) PosSupplierController {
	return &posSupplierController{
		service: service,
	}
}
