package auth

import (
	"go-adv/4-order-api/pkg/req"
	"go-adv/4-order-api/pkg/res"
	"net/http"
)

type AuthHandler struct {
	AuthService *AuthService
}

type AuthHandlerDeps struct {
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{AuthService: deps.AuthService}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/verify", handler.Verify())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[UserCreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessionId, err := handler.AuthService.Login(body.PhoneNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := &LoginResponse{
			SessionId: sessionId,
		}
		res.Json(w, 201, data)
	}
}

func (handler *AuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[UserVerifyRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := handler.AuthService.Verify(body.SessionId, body.Code)
		if err != nil || token == "" {
			http.Error(w, ERR_UNAUTHORIZED, http.StatusUnauthorized)
			return
		}

		data := &VerifyResponse{
			Token: token,
		}
		res.Json(w, 200, data)
	}
}
