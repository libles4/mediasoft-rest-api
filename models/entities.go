package models

type Car struct {
	ID          int    `json:"id"`
	BrandName   string `json:"brand_name"`
	ModelName   string `json:"model_name"`
	Mileage     int    `json:"mileage"`
	OwnersCount int    `json:"owners_count"`
}

type Flower struct {
	ID           int     `json:"id"`
	FlowerName   string  `json:"flower_name"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	DeliveryDate string  `json:"delivery_date"`
}

type Furniture struct {
	ID            int    `json:"id"`
	FurnitureName string `json:"furniture_name"`
	Manufacturer  string `json:"manufacturer"`
	Height        int    `json:"height"`
	Width         int    `json:"width"`
	Length        int    `json:"length"`
}
