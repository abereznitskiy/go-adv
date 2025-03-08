package main

import (
	"fmt"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/internal/product"
	"go-adv/4-order-api/pkg/customValidate"
	"go-adv/4-order-api/pkg/db"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	productRepository := product.NewProductRepository(db)
	product.NewProductHandler(router, product.ProductHandlerDeps{ProductRepository: productRepository})
	validate := validator.New()
	validate.RegisterValidation("string_array", customValidate.StringArrayValidation)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening 8081")
	server.ListenAndServe()
}
