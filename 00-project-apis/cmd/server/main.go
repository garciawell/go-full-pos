package main

import (
	"fmt"
	"net/http"

	"github.com/garciawell/go-full-pos/apis/configs"
	_ "github.com/garciawell/go-full-pos/apis/docs"
	"github.com/garciawell/go-full-pos/apis/internal/entity"
	"github.com/garciawell/go-full-pos/apis/internal/infra/database"
	"github.com/garciawell/go-full-pos/apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title          Go Experts API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
var productHandler *handlers.ProductHandler
var userHandler *handlers.UserHandler
var conf *configs.Conf

func init() {
	config, err := configs.LoadConfig(".")
	conf = config
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	productHandler = handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler = handlers.NewUserHandler(userDB, config.JwtExpiresIn)
}

func main() {
	r := chi.NewRouter()
	// r.Use(LogRequest)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", conf.TokenAuthKey))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(conf.TokenAuthKey))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("//{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate-token", userHandler.GetJWT)

	fmt.Println("Server is running on port 8000...")

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)

}

// func LogRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }
