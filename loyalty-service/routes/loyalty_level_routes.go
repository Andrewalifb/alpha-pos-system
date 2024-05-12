package routes

import (
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/controller"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosLoyaltyLevelRoutes(r *gin.Engine, posLoyaltyLevelController controller.PosLoyaltyLevelController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(midlleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1/loyalty-levels")
	// Create New PosLoyaltyLevel
	routesV1.POST("/pos_loyalty_level", posLoyaltyLevelController.HandleCreatePosLoyaltyLevelRequest)
	// Get PosLoyaltyLevel by ID
	routesV1.GET("/pos_loyalty_level/:id", posLoyaltyLevelController.HandleReadPosLoyaltyLevelRequest)
	// Update Existing PosLoyaltyLevel
	routesV1.PUT("/pos_loyalty_level/:id", posLoyaltyLevelController.HandleUpdatePosLoyaltyLevelRequest)
	// Delete PosLoyaltyLevel
	routesV1.DELETE("/pos_loyalty_level/:id", posLoyaltyLevelController.HandleDeletePosLoyaltyLevelRequest)
	// Get All PosLoyaltyLevels
	routesV1.GET("/pos_loyalty_levels", posLoyaltyLevelController.HandleReadAllPosLoyaltyLevelsRequest)
}
