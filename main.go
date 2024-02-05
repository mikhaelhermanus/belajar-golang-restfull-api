package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	productController "belajar-golang-restful-api/controller/products"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	productRepository "belajar-golang-restful-api/repository/products"
	"belajar-golang-restful-api/service"
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

	router := app.NewRouter(categoryController, productController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	log.Println("running on port : ", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
