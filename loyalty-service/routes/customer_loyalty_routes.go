package routes

import (
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/controller"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosCustomerLoyaltyRoutes(r *gin.Engine, posCustomerLoyaltyController controller.PosCustomerLoyaltyController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(midlleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1/customer-loyalties")
	// Create New PosCustomerLoyalty
	routesV1.POST("/pos_customer_loyalty", posCustomerLoyaltyController.HandleCreatePosCustomerLoyaltyRequest)
	// Get PosCustomerLoyalty by ID
	routesV1.GET("/pos_customer_loyalty/:id", posCustomerLoyaltyController.HandleReadPosCustomerLoyaltyRequest)
	// Update Existing PosCustomerLoyalty
	routesV1.PUT("/pos_customer_loyalty/:id", posCustomerLoyaltyController.HandleUpdatePosCustomerLoyaltyRequest)
	// Delete PosCustomerLoyalty
	routesV1.DELETE("/pos_customer_loyalty/:id", posCustomerLoyaltyController.HandleDeletePosCustomerLoyaltyRequest)
	// Get All PosCustomerLoyalties
	routesV1.GET("/pos_customer_loyalties", posCustomerLoyaltyController.HandleReadAllPosCustomerLoyaltiesRequest)
}
