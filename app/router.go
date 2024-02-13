package app

import (
	"belajar-golang-restful-api/controller"
	productController "belajar-golang-restful-api/controller/products"
	userController "belajar-golang-restful-api/controller/users"
	"belajar-golang-restful-api/exception"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

// ValidateMiddleware validates the JWT token.
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
	})
}

func NewRouter(categoryController controller.CategoryController, productController productController.ProductController) *httprouter.Router {
	router := httprouter.New()

	// router.GET("/api/categories", categoryController.FindAll)
	// router.GET("/api/categories/:categoryId", categoryController.FindById)
	// router.POST("/api/categories", categoryController.Create)
	// router.PUT("/api/categories/:categoryId", categoryController.Update)
	// router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	// products
	// router.POST("/api/products", productController.Create)
	// router.GET("/api/products", productController.FindAll)
	// router.GET("/api/products/:productId", productController.FindById)
	// router.PUT("/api/products/:productId", productController.Update)
	// router.DELETE("/api/products/:productId", productController.Delete)
	// router.POST("/api/login", userController.CreateToken)
	router.PanicHandler = exception.ErrorHandler

	return router
}

func MuxRouter(categoryController controller.CategoryController, productController productController.ProductController, userController userController.UserController) *mux.Router {
	routerMux := mux.NewRouter()
	// Category Service
	routerMux.HandleFunc("/api/categories", categoryController.FindAll).Methods("Get")
	routerMux.HandleFunc("/api/categories/{categoryId}", categoryController.FindById).Methods("Get")
	routerMux.HandleFunc("/api/categories", categoryController.Create).Methods("Post")
	routerMux.HandleFunc("/api/categories/{categoryId}", categoryController.Update).Methods("Put")
	routerMux.HandleFunc("/api/categories/{categoryId}", categoryController.Delete).Methods("Delete")
	//product service
	routerMux.HandleFunc("/api/products", productController.Create).Methods("Post")
	routerMux.HandleFunc("/api/products", productController.FindAll).Methods("Get")
	routerMux.HandleFunc("/api/products/{productId}", productController.FindById).Methods("Get")
	routerMux.HandleFunc("/api/products/{productId}", productController.Update).Methods("PUT")
	routerMux.HandleFunc("/api/products/{productId}", productController.Delete).Methods("Delete")
	//user service
	// routerMux.HandleFunc("/api/login", userController.CreateToken).Methods("Post")
	routerMux.HandleFunc("/api/user/register", userController.CreateUser).Methods("post")

	return routerMux
}
