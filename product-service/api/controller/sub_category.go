package controller

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/gin-gonic/gin"
)

type PosProductSubCategoryController interface {
	HandleCreatePosProductSubCategoryRequest(c *gin.Context)
	HandleReadPosProductSubCategoryRequest(c *gin.Context)
	HandleUpdatePosProductSubCategoryRequest(c *gin.Context)
	HandleDeletePosProductSubCategoryRequest(c *gin.Context)
	HandleReadAllPosProductSubCategoriesRequest(c *gin.Context)
}

type posProductSubCategoryController struct {
	service pb.PosProductSubCategoryServiceClient
}

func NewPosProductSubCategoryController(service pb.PosProductSubCategoryServiceClient) PosProductSubCategoryController {
	return &posProductSubCategoryController{
		service: service,
	}
}
