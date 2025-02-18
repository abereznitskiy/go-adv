package main

import (
	"fmt"
	"net/http"
)

func main() {
	PORT := "3000"
	router := http.NewServeMux()
	NewDiceHandler(router)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: router,
	}
	fmt.Printf("Server started on PORT: %v", PORT)
	server.ListenAndServe()
}
