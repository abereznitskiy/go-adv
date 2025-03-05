package req

import (
	"go-adv/3-validation-api/pkg/res"

	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := DecodeBody[T](r.Body)
	if err != nil {
		res.Json(*w, 402, err.Error())
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		res.Json(*w, 402, err.Error())
		return nil, err
	}

	return &body, nil
}
