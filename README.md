# REST API example application

This is a simple products, category and oder product example of a personal project to made a REST
API with GO.

## Setup
edit `env.example` for your local user and databse to use
create any database and table you need.

## run the app
`go run main.go`

# REST API

The REST API to the example app is described below.

## Request
CATEGORY API

`GET /api/categories` Get all category of product </br>
`POST /api/categories` Create Category of product </br>
`GET /api/categories/{id}` Get Category by Id </br>
`PUT /api/categories/{id}` Update category by id </br>
`DELETE /api/categories/{id}` Delete Category By id </br>

PRODUCT API

`GET /api/products` Get all product </br>
`POST /api/products` Create product </br>
`GET /api/products/{id}` Get products by Id </br>
`PUT /api/products/{id}` Update products by id </br>
`DELETE /api/products/{id}` Delete products By id </br>

AUTH API

`POST api/user/register` Create New user API </br>
`POST /api/login` Login User </br>

Order API
`POST api/order/create` Create Order </br>
`GET /order-detail/{order_id} Get Order By id` </br>

## Room of Improvement
- Make an API for create AUTH base on role
- Make an middleware for every api need base on role
- Create A wallet API with one to manny relation database
- create Stock management of product

