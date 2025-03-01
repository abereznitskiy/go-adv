package verify

import (
	"fmt"
	"go-adv/3-validation-api/configs"
	"go-adv/3-validation-api/pkg/email"
	"go-adv/3-validation-api/pkg/files"
	"go-adv/3-validation-api/pkg/hash"
	"go-adv/3-validation-api/pkg/req"
	"go-adv/3-validation-api/pkg/res"

	"net/http"
)

type VerifyHandlerDeps struct {
	Config *configs.Config
	Db     *files.JsonDb
}

type VerifyHandler struct {
	Config *configs.Config
	Db     *files.JsonDb
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{Config: deps.Config, Db: deps.Db}
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
	router.HandleFunc("POST /send", handler.Send())
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		_, exists := handler.Db.Get(hash)
		if !exists {
			res.Json(w, 402, "Error")
			return
		}

		err := handler.Db.Delete(hash)
		if err != nil {
			res.Json(w, 500, "Error deleting from db")
			return
		}
		res.Json(w, 200, "Email verified")
	}
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}

		hash := hash.EncodeEmail(body.Email)
		link := fmt.Sprintf("%s%s:%s/%s", configs.PROTOCOL, configs.DOMAIN, configs.PORT, hash)
		err = email.Send(email.SendParams{
			UserEmail:      body.Email,
			Link:           link,
			SmtpUsername:   handler.Config.Email,
			SmtpPassword:   handler.Config.Password,
			ResponseWriter: &w,
		})
		if err != nil {
			return
		}

		err = handler.Db.Set(hash, body.Email)
		if err != nil {
			res.Json(w, 500, "Error setting to DB")
			return
		}

		successMessage := "Link sended to email - " + body.Email
		res.Json(w, 200, successMessage)
	}
}
