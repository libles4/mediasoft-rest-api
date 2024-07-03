package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/libles4/mediasoft-rest-api/models"
	"github.com/libles4/mediasoft-rest-api/utils"
)

var flowerMutex sync.Mutex

func GetFlowers(context *gin.Context) {
	flowerMutex.Lock()
	defer flowerMutex.Unlock()
	context.JSON(http.StatusOK, utils.Data.FlowerBase)
}

func CreateFlower(c *gin.Context) {
	flowerMutex.Lock()
	defer flowerMutex.Unlock()
	var item models.Flower
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = len(utils.Data.FlowerBase) + 1
	utils.Data.FlowerBase = append(utils.Data.FlowerBase, item)
	utils.SaveData()
	c.JSON(http.StatusCreated, item)
}

func GetFlowerByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	flowerMutex.Lock()
	defer flowerMutex.Unlock()
	for _, item := range utils.Data.FlowerBase {
		if item.ID == id {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
}

func UpdateFlowerByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	flowerMutex.Lock()
	defer flowerMutex.Unlock()
	for i, item := range utils.Data.FlowerBase {
		if item.ID == id {
			if err := context.BindJSON(&item); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			item.ID = id
			utils.Data.FlowerBase[i] = item
			utils.SaveData()
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
}

func DeleteFlowerByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	flowerMutex.Lock()
	defer flowerMutex.Unlock()
	for i, item := range utils.Data.FlowerBase {
		if item.ID == id {
			utils.Data.FlowerBase = append(utils.Data.FlowerBase[:i], utils.Data.FlowerBase[i+1:]...)
			utils.SaveData()
			context.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
}
