package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosInvoiceController interface {
	HandleCreatePosInvoiceRequest(c *gin.Context)
	HandleReadPosInvoiceRequest(c *gin.Context)
	HandleUpdatePosInvoiceRequest(c *gin.Context)
	HandleDeletePosInvoiceRequest(c *gin.Context)
	HandleReadAllPosInvoicesRequest(c *gin.Context)
}

type posInvoiceController struct {
	service pb.PosInvoiceServiceClient
}

func NewPosInvoiceController(service pb.PosInvoiceServiceClient) PosInvoiceController {
	return &posInvoiceController{
		service: service,
	}
}
