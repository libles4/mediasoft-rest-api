package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/libles4/mediasoft-rest-api/handlers"
	"github.com/libles4/mediasoft-rest-api/utils"
)

// Определяем структуры для данных
func main() {
	router := gin.Default()

	router.GET("/cars", handlers.GetCars)
	router.POST("/cars", handlers.CreateCar)
	router.GET("/cars/:id", handlers.GetCarByID)
	router.PUT("/cars/:id", handlers.UpdateCarByID)
	router.DELETE("/cars/:id", handlers.DeleteCarByID)

	router.GET("/furniture", handlers.GetFurniture)
	router.POST("/furniture", handlers.CreateFurniture)
	router.GET("/furniture/:id", handlers.GetFurnitureByID)
	router.PUT("/furniture/:id", handlers.UpdateFurnitureByID)
	router.DELETE("/furniture/:id", handlers.DeleteFurnitureByID)

	router.GET("/flowers", handlers.GetFlowers)
	router.POST("/flowers", handlers.CreateFlower)
	router.GET("/flowers/:id", handlers.GetFlowerByID)
	router.PUT("/flowers/:id", handlers.UpdateFlowerByID)
	router.DELETE("/flowers/:id", handlers.DeleteFlowerByID)

	utils.LoadData()

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
