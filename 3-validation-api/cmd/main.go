package main

import (
	"fmt"
	"go-adv/3-validation-api/configs"
	"go-adv/3-validation-api/internal/verify"
	"go-adv/3-validation-api/pkg/files"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{Config: conf, Db: files.NewJsonDb(conf.JsonDbPath)})

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", conf.Port),
		Handler: router,
	}
	fmt.Println("Server is listening 8081")
	server.ListenAndServe()
}
