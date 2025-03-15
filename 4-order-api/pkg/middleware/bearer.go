package middleware

import (
	"context"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/pkg/jwt"

	"net/http"
	"strings"
)

type key string

const (
	CONTEXT_PHONE_NUMBER_KEY key    = "ContextPhoneNumberKey"
	AUTHORIZATION_KEY        string = "Authorization"
	BEARER_PREFIX            string = "Bearer "
)

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(AUTHORIZATION_KEY)
		if !strings.HasPrefix(authHeader, BEARER_PREFIX) {
			writeUnauthed(w)
			return
		}

		token := strings.TrimPrefix(authHeader, BEARER_PREFIX)
		isValid, data := jwt.NewJWT(config.Db.Secret).Parse(token)
		if !isValid {
			writeUnauthed(w)
			return
		}

		ctx := context.WithValue(r.Context(), CONTEXT_PHONE_NUMBER_KEY, data.PhoneNumber)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}
