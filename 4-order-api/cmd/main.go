package main

import (
	"fmt"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening 8081")
	server.ListenAndServe()
}
