package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/libles4/mediasoft-rest-api/models"
	"github.com/libles4/mediasoft-rest-api/utils"
)

var furnitureMutex sync.Mutex

func GetFurniture(c *gin.Context) {
	furnitureMutex.Lock()
	defer furnitureMutex.Unlock()
	c.JSON(http.StatusOK, utils.Data.Furniture)
}

func CreateFurniture(c *gin.Context) {
	furnitureMutex.Lock()
	defer furnitureMutex.Unlock()
	var item models.Furniture
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = len(utils.Data.Furniture) + 1
	utils.Data.Furniture = append(utils.Data.Furniture, item)
	utils.SaveData()
	c.JSON(http.StatusCreated, item)
}

func GetFurnitureByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	furnitureMutex.Lock()
	defer furnitureMutex.Unlock()
	for _, item := range utils.Data.Furniture {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Furniture not found"})
}

func UpdateFurnitureByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	furnitureMutex.Lock()
	defer furnitureMutex.Unlock()
	for i, item := range utils.Data.Furniture {
		if item.ID == id {
			if err := c.BindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			item.ID = id
			utils.Data.Furniture[i] = item
			utils.SaveData()
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Furniture not found"})
}

func DeleteFurnitureByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	furnitureMutex.Lock()
	defer furnitureMutex.Unlock()
	for i, item := range utils.Data.Furniture {
		if item.ID == id {
			utils.Data.Furniture = append(utils.Data.Furniture[:i], utils.Data.Furniture[i+1:]...)
			utils.SaveData()
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Furniture not found"})
}
