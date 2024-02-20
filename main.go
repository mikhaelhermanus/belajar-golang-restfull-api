package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	orderController "belajar-golang-restful-api/controller/orders"
	productController "belajar-golang-restful-api/controller/products"
	userController "belajar-golang-restful-api/controller/users"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/repository"
	userRepository "belajar-golang-restful-api/repository/auth"
	orderRepository "belajar-golang-restful-api/repository/orders"
	productRepository "belajar-golang-restful-api/repository/products"
	"belajar-golang-restful-api/service"
	userService "belajar-golang-restful-api/service/auth"
	orderService "belajar-golang-restful-api/service/orders"
	productService "belajar-golang-restful-api/service/products"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	local := goDotEnvVariable("USER_LOCAL")
	database := goDotEnvVariable("DATABASE_LOCAL")
	db := app.NewDB(local, database)
	validate := validator.New()
	//category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	//product
	productRepository := productRepository.NewProductsRepository()
	productService := productService.NewProductService(productRepository, db, validate)
	productController := productController.NewProductController(productService)
	// user
	userRepository := userRepository.NewAuthsRepository()
	userService := userService.NewAuthService(userRepository, db, validate)
	userController := userController.NewUserController(userService)
	// order
	orderRepository := orderRepository.NewOrdersRepository()
	orderService := orderService.NewOrderService(orderRepository, db, validate)
	orderController := orderController.NewOrderController(orderService)
	// userController := userController.
	// router := app.NewRouter(categoryController, productController)
	router := app.MuxRouter(categoryController, productController, userController, orderController)
	// routerLogin := app.NewRouter(routerlogin)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	// mux router version
	// mux := mux.NewRouter()
	// mux.HandleFunc("/api/categories", productController.FindAll).Methods("Get")

	log.Println("running on port : ", server.Addr)
	// log.Fatal(http.ListenAndServe(":3000", mux))
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
