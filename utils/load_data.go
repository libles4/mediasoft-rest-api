package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/libles4/mediasoft-rest-api/models"
)

type DataStore struct {
	Cars       []models.Car       `json:"cars"`
	Furniture  []models.Furniture `json:"furniture"`
	FlowerBase []models.Flower    `json:"flower_base"`
}

var Data DataStore

func LoadData() {
	file, err := os.Open("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			saveInitialData()
			return
		}
		log.Fatalf("Failed to open data file: %v", err)
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &Data)
}

func SaveData() {
	bytes, err := json.MarshalIndent(Data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}
	err = ioutil.WriteFile("data.json", bytes, 0644)
	if err != nil {
		log.Fatalf("Failed to save data file: %v", err)
	}
}

func saveInitialData() {
	Data = DataStore{
		Cars:       []models.Car{},
		Furniture:  []models.Furniture{},
		FlowerBase: []models.Flower{},
	}
	SaveData()
}
