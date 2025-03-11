package main

import (
	"fmt"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/internal/product"
	"go-adv/4-order-api/pkg/customValidate"
	"go-adv/4-order-api/pkg/db"
	"go-adv/4-order-api/pkg/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	conf := configs.LoadConfig()
	db, err := db.NewDb(conf)
	if err != nil {
		fmt.Printf("Database error: %s", err)
	}
	router := http.NewServeMux()
	productRepository := product.NewProductRepository(db)
	product.NewProductHandler(router, product.ProductHandlerDeps{ProductRepository: productRepository})
	validate := validator.New()
	validate.RegisterValidation("string_array", customValidate.StringArrayValidation)

	stack := middleware.Chain(middleware.Log)
	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("Server is listening 8081")
	server.ListenAndServe()
}
