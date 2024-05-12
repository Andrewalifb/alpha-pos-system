package routes

import (
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/controller"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosRewardRoutes(r *gin.Engine, posRewardController controller.PosRewardController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(midlleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1/rewards")
	// Create New PosReward
	routesV1.POST("/pos_reward", posRewardController.HandleCreatePosRewardRequest)
	// Get PosReward by ID
	routesV1.GET("/pos_reward/:id", posRewardController.HandleReadPosRewardRequest)
	// Update Existing PosReward
	routesV1.PUT("/pos_reward/:id", posRewardController.HandleUpdatePosRewardRequest)
	// Delete PosReward
	routesV1.DELETE("/pos_reward/:id", posRewardController.HandleDeletePosRewardRequest)
	// Get All PosRewards
	routesV1.GET("/pos_rewards", posRewardController.HandleReadAllPosRewardsRequest)
}
