package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosCompanyController interface {
	HandleCreatePosCompanyRequest(c *gin.Context)
	HandleReadPosCompanyRequest(c *gin.Context)
	HandleUpdatePosCompanyRequest(c *gin.Context)
	HandleDeletePosCompanyRequest(c *gin.Context)
	HandleReadAllPosCompaniesRequest(c *gin.Context)
}

type posCompanyController struct {
	service pb.PosCompanyServiceClient
}

func NewPosCompanyController(service pb.PosCompanyServiceClient) PosCompanyController {
	return &posCompanyController{
		service: service,
	}
}
