package order

import (
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/pkg/middleware"
	"go-adv/4-order-api/pkg/models"
	"go-adv/4-order-api/pkg/req"
	"go-adv/4-order-api/pkg/res"
	"net/http"
	"strconv"
)

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	Config          *configs.Config
}

type OrderHandler struct {
	OrderRepository *OrderRepository
	Config          *configs.Config
}

func NewProductHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{OrderRepository: deps.OrderRepository}

	router.Handle("POST /order", middleware.IsAuth(handler.Create(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuth(handler.GetById(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuth(handler.GetMyOrders(), deps.Config))

}

func (handler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authData := r.Context().Value(middleware.CONTEXT_AUTH_DATA_KEY).(*middleware.AuthContext)
		body, err := req.HandleBody[OrderCreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		parsedId, err := strconv.ParseUint(authData.UserId, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		order := &models.Order{
			UserID:   uint(parsedId),
			Products: body.Products,
		}
		result, err := handler.OrderRepository.Create(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, 201, result.ID)
	}
}

func (handler *OrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathId := r.PathValue("id")
		authData := r.Context().Value(middleware.CONTEXT_AUTH_DATA_KEY).(*middleware.AuthContext)
		parsedId, err := strconv.ParseUint(authData.UserId, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := handler.OrderRepository.GetById(pathId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if result.UserID != uint(parsedId) {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		res.Json(w, 200, result)
	}
}

func (handler *OrderHandler) GetMyOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authData := r.Context().Value(middleware.CONTEXT_AUTH_DATA_KEY).(*middleware.AuthContext)

		result, err := handler.OrderRepository.GetAllByUserId(authData.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, 200, result)
	}
}
