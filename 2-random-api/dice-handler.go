package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type DiceHandler struct{}

func NewDiceHandler(router *http.ServeMux) {
	handler := &DiceHandler{}
	router.HandleFunc("/dice", handler.dice())
}

func (handler *DiceHandler) dice() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprint(rand.Intn(6) + 1)))
	}
}
