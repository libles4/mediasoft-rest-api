package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/libles4/mediasoft-rest-api/models"
	"github.com/libles4/mediasoft-rest-api/utils"
)

var carMutex sync.Mutex

func GetCars(c *gin.Context) {
	carMutex.Lock()
	defer carMutex.Unlock()
	c.JSON(http.StatusOK, utils.Data.Cars)
}

func CreateCar(c *gin.Context) {
	carMutex.Lock()
	defer carMutex.Unlock()
	var car models.Car
	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	car.ID = len(utils.Data.Cars) + 1
	utils.Data.Cars = append(utils.Data.Cars, car)
	utils.SaveData()
	c.JSON(http.StatusCreated, car)
}

func GetCarByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	carMutex.Lock()
	defer carMutex.Unlock()
	for _, car := range utils.Data.Cars {
		if car.ID == id {
			c.JSON(http.StatusOK, car)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
}

func UpdateCarByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	carMutex.Lock()
	defer carMutex.Unlock()
	for i, car := range utils.Data.Cars {
		if car.ID == id {
			if err := c.BindJSON(&car); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			car.ID = id
			utils.Data.Cars[i] = car
			utils.SaveData()
			c.JSON(http.StatusOK, car)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
}

func DeleteCarByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	carMutex.Lock()
	defer carMutex.Unlock()
	for i, car := range utils.Data.Cars {
		if car.ID == id {
			utils.Data.Cars = append(utils.Data.Cars[:i], utils.Data.Cars[i+1:]...)
			utils.SaveData()
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
}
