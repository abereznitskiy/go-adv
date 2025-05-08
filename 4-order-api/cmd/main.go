package main

import (
	"fmt"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/internal/auth"
	"go-adv/4-order-api/internal/order"
	"go-adv/4-order-api/internal/product"
	"go-adv/4-order-api/internal/user"
	"go-adv/4-order-api/pkg/customValidate"
	"go-adv/4-order-api/pkg/db"
	"go-adv/4-order-api/pkg/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func App() http.Handler {
	log.SetFormatter(&log.JSONFormatter{})

	conf := configs.LoadConfig()
	db, err := db.NewDb(conf)
	if err != nil {
		fmt.Printf("Database error: %s", err)
	}
	router := http.NewServeMux()
	productRepository := product.NewProductRepository(db)
	userRepository := user.NewUserRepository(db)
	orderRepository := order.NewOrderRepository(db)
	authService := auth.NewAuthService(userRepository, conf)
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
		Config:            conf,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{AuthService: authService})
	order.NewProductHandler(router, order.OrderHandlerDeps{
		OrderRepository: orderRepository,
		Config:          conf,
	})
	validate := validator.New()
	validate.RegisterValidation("string_array", customValidate.StringArrayValidation)

	stack := middleware.Chain(middleware.Log)

	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}
	fmt.Println("Server is listening 8081")
	server.ListenAndServe()
}
