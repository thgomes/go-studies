package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/thgomes/go-studies/5-first-api/configs"
	_ "github.com/thgomes/go-studies/5-first-api/docs"
	"github.com/thgomes/go-studies/5-first-api/internal/entity"
	"github.com/thgomes/go-studies/5-first-api/internal/infra/database"
	"github.com/thgomes/go-studies/5-first-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title First Go API Example
// @version 1.0
// @description My first rest API in Go

// @contact.name Thiago Gomes
// @contact.email thiagosiqgomes@gmail.com

// @licence.name License Example
// @licence.url http://wwww.license.example.com

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.JWTExperesIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExperesIn", configs.JWTExperesIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
