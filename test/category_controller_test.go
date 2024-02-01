package test

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler {
	db := setupTestDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

// go test -v -run TestCreateCategorySuccess
func TestCreateCategorySuccess(t *testing.T) {
	// db := setupTestDB()
	// truncateCategory(db)
	router := setupRouter()

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	// body, _ := io.ReadAll(response.Body)
	// var responseBody map[string]interface{}
	// json.Unmarshal(body, &responseBody)

	// assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// assert.Equal(t, "OK", responseBody["status"])
	// assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	// db := setupTestDB()
	// truncateCategory(db)
	// router := setupRouter(db)

	// requestBody := strings.NewReader(`{"name" : ""}`)
	// request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	// request.Header.Add("Content-Type", "application/json")
	// request.Header.Add("X-API-Key", "RAHASIA")

	// recorder := httptest.NewRecorder()

	// router.ServeHTTP(recorder, request)

	// response := recorder.Result()
	// assert.Equal(t, 400, response.StatusCode)

	// body, _ := io.ReadAll(response.Body)
	// var responseBody map[string]interface{}
	// json.Unmarshal(body, &responseBody)

	// assert.Equal(t, 400, int(responseBody["code"].(float64)))
	// assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	// db := setupTestDB()
	// truncateCategory(db)

	// tx, _ := db.Begin()
	// categoryRepository := repository.NewCategoryRepository()
	// category := categoryRepository.Save(context.Background(), tx, domain.Category{
	// 	Name: "Gadget",
	// })
	// tx.Commit()

	// router := setupRouter(db)

	// requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	// request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	// request.Header.Add("Content-Type", "application/json")
	// request.Header.Add("X-API-Key", "RAHASIA")

	// recorder := httptest.NewRecorder()

	// router.ServeHTTP(recorder, request)

	// response := recorder.Result()
	// assert.Equal(t, 200, response.StatusCode)

	// body, _ := io.ReadAll(response.Body)
	// var responseBody map[string]interface{}
	// json.Unmarshal(body, &responseBody)

	// assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// assert.Equal(t, "OK", responseBody["status"])
	// assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	// db := setupTestDB()
	// truncateCategory(db)

	// tx, _ := db.Begin()
	// categoryRepository := repository.NewCategoryRepository()
	// category := categoryRepository.Save(context.Background(), tx, domain.Category{
	// 	Name: "Gadget",
	// })
	// tx.Commit()

	// router := setupRouter(db)

	// requestBody := strings.NewReader(`{"name" : ""}`)
	// request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	// request.Header.Add("Content-Type", "application/json")
	// request.Header.Add("X-API-Key", "RAHASIA")

	// recorder := httptest.NewRecorder()

	// router.ServeHTTP(recorder, request)

	// response := recorder.Result()
	// assert.Equal(t, 400, response.StatusCode)

	// body, _ := io.ReadAll(response.Body)
	// var responseBody map[string]interface{}
	// json.Unmarshal(body, &responseBody)

	// assert.Equal(t, 400, int(responseBody["code"].(float64)))
	// assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
