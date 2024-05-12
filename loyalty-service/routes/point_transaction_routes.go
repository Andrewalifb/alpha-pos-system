package routes

import (
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/controller"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosPointsTransactionRoutes(r *gin.Engine, posPointsTransactionController controller.PosPointsTransactionController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(midlleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1/points-transactions")
	// Create New PosPointsTransaction
	routesV1.POST("/pos_points_transaction", posPointsTransactionController.HandleCreatePosPointsTransactionRequest)
	// Get PosPointsTransaction by ID
	routesV1.GET("/pos_points_transaction/:id", posPointsTransactionController.HandleReadPosPointsTransactionRequest)
	// Update Existing PosPointsTransaction
	routesV1.PUT("/pos_points_transaction/:id", posPointsTransactionController.HandleUpdatePosPointsTransactionRequest)
	// Delete PosPointsTransaction
	routesV1.DELETE("/pos_points_transaction/:id", posPointsTransactionController.HandleDeletePosPointsTransactionRequest)
	// Get All PosPointsTransactions
	routesV1.GET("/pos_points_transactions", posPointsTransactionController.HandleReadAllPosPointsTransactionsRequest)
}
