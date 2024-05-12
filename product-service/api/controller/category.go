package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosProductCategoryController interface {
	HandleCreatePosProductCategoryRequest(c *gin.Context)
	HandleReadPosProductCategoryRequest(c *gin.Context)
	HandleUpdatePosProductCategoryRequest(c *gin.Context)
	HandleDeletePosProductCategoryRequest(c *gin.Context)
	HandleReadAllPosProductCategoriesRequest(c *gin.Context)
}

type posProductCategoryController struct {
	service pb.PosProductCategoryServiceClient
}

func NewPosProductCategoryController(service pb.PosProductCategoryServiceClient) PosProductCategoryController {
	return &posProductCategoryController{
		service: service,
	}
}
