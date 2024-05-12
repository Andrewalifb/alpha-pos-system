package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosStoreBranchController interface {
	HandleCreatePosStoreBranchRequest(c *gin.Context)
	HandleReadPosStoreBranchRequest(c *gin.Context)
	HandleUpdatePosStoreBranchRequest(c *gin.Context)
	HandleDeletePosStoreBranchRequest(c *gin.Context)
	HandleReadAllPosStoreBranchesRequest(c *gin.Context)
}

type posStoreBranchController struct {
	service pb.PosStoreBranchServiceClient
}

func NewPosStoreBranchController(service pb.PosStoreBranchServiceClient) PosStoreBranchController {
	return &posStoreBranchController{
		service: service,
	}
}
