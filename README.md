# Практика в MediaSoft

Вот примеры JSON-ответов для `GET` запросов к каждой сущности:

1. GET /cars:

[
  {
    "id": 1,
    "brand_name": "Toyota",
    "model_name": "Camry",
    "mileage": 100000,
    "owners_count": 2
  },
  {
    "id": 2,
    "brand_name": "Honda",
    "model_name": "Civic",
    "mileage": 50000,
    "owners_count": 1
  }
]

Получение данных конкретного автомобиля по id

2. GET /cars/1

{
  "id": 1,
  "brand_name": "Toyota",
  "model_name": "Camry",
  "mileage": 100000,
  "owners_count": 2
}


3. GET /furniture:

[
  {
    "id": 1,
    "furniture_name": "Sofa",
    "manufacturer": "IKEA",
    "height": 80,
    "width": 200,
    "length": 150
  },
  {
    "id": 2,
    "furniture_name": "Table",
    "manufacturer": "Ashley",
    "height": 75,
    "width": 120,
    "length": 180
  }
]


4. GET /furniture/1

{
  "id": 1,
  "furniture_name": "Sofa",
  "manufacturer": "IKEA",
  "height": 80,
  "width": 200,
  "length": 150
}


5. GET /flowers

[
  {
    "id": 1,
    "flower_name": "Rose",
    "quantity": 10,
    "price": 15.99,
    "delivery_date": "2023-12-10"
  },
  {
    "id": 2,
    "flower_name": "Lily",
    "quantity": 5,
    "price": 12.50,
    "delivery_date": "2023-12-15"
  }
]


6. GET /flowers/1

{
  "id": 1,
  "flower_name": "Rose",
  "quantity": 10,
  "price": 15.99,
  "delivery_date": "2023-12-10"
}

примеры POST-запросов к каждой сущности, Примеры JSON-ответов:

1. POST /cars

Запрос:

{
  "brand_name": "Ford",
  "model_name": "Mustang",
  "mileage": 30000,
  "owners_count": 1
}


Ответ:

{
  "id": 3,
  "brand_name": "Ford",
  "model_name": "Mustang",
  "mileage": 30000,
  "owners_count": 1
}


2. POST /furniture

Запрос:

{
  "furniture_name": "Chair",
  "manufacturer": "Herman Miller",
  "height": 85,
  "width": 60,
  "length": 60
}


Ответ:

{
  "id": 3,
  "furniture_name": "Chair",
  "manufacturer": "Herman Miller",
  "height": 85,
  "width": 60,
  "length": 60
}


3. POST /flowers

Запрос:

{
  "flower_name": "Tulip",
  "quantity": 20,
  "price": 10.00,
  "delivery_date": "2024-01-15"
}


Ответ:

{
  "id": 3,
  "flower_name": "Tulip",
  "quantity": 20,
  "price": 10.00,
  "delivery_date": "2024-01-15"
}


