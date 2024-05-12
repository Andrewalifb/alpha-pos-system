package routes

import (
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/controller"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosRewardRedemptionRoutes(r *gin.Engine, posRewardRedemptionController controller.PosRewardRedemptionController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(midlleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1/reward-redemptions")
	// Create New PosRewardRedemption
	routesV1.POST("/pos_reward_redemption", posRewardRedemptionController.HandleCreatePosRewardRedemptionRequest)
	// Get PosRewardRedemption by ID
	routesV1.GET("/pos_reward_redemption/:id", posRewardRedemptionController.HandleReadPosRewardRedemptionRequest)
	// Update Existing PosRewardRedemption
	routesV1.PUT("/pos_reward_redemption/:id", posRewardRedemptionController.HandleUpdatePosRewardRedemptionRequest)
	// Delete PosRewardRedemption
	routesV1.DELETE("/pos_reward_redemption/:id", posRewardRedemptionController.HandleDeletePosRewardRedemptionRequest)
	// Get All PosRewardRedemptions
	routesV1.GET("/pos_reward_redemptions", posRewardRedemptionController.HandleReadAllPosRewardRedemptionsRequest)
}
