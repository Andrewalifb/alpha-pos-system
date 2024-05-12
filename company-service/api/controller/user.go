package controller

import (
	"net/http"

	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system/company-service/utils"
	"github.com/gin-gonic/gin"
)

type PosUserController interface {
	HandleCreatePosUserRequest(c *gin.Context)
	HandleReadPosUserRequest(c *gin.Context)
	HandleUpdatePosUserRequest(c *gin.Context)
	HandleDeletePosUserRequest(c *gin.Context)
	HandleReadAllPosUsersRequest(c *gin.Context)
	HandleLoginRequest(c *gin.Context)
}

type posUserController struct {
	service pb.PosUserServiceClient
}

func NewPosUserController(service pb.PosUserServiceClient) PosUserController {
	return &posUserController{
		service: service,
	}
}

func (c *posUserController) HandleCreatePosUserRequest(ctx *gin.Context) {
	var req pb.CreatePosUserRequest
	if err := ctx.ShouldBindJSON(&req.PosUser); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	res, err := c.service.CreatePosUser(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_USER, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posUserController) HandleLoginRequest(ctx *gin.Context) {
	var req pb.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	res, err := c.service.Login(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess("Login successful", res)
	ctx.JSON(http.StatusOK, successResponse)
}
