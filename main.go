package main

import (
	"githhub.com/gign-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-resturant-management/database"
	"golang-resturant-management/middleware"
	"golang-resturant-management/routes"
	"os"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
