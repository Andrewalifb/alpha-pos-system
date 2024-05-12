package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosRoleController interface {
	HandleCreatePosRoleRequest(c *gin.Context)
	HandleReadPosRoleRequest(c *gin.Context)
	HandleUpdatePosRoleRequest(c *gin.Context)
	HandleDeletePosRoleRequest(c *gin.Context)
	HandleReadAllPosRolesRequest(c *gin.Context)
}

type posRoleController struct {
	service pb.PosRoleServiceClient
}

func NewPosRoleController(service pb.PosRoleServiceClient) PosRoleController {
	return &posRoleController{
		service: service,
	}
}
