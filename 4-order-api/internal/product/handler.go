package product

import (
	"fmt"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/pkg/middleware"
	"go-adv/4-order-api/pkg/req"
	"go-adv/4-order-api/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
	Config            *configs.Config
}

type ProductHandler struct {
	ProductRepository *ProductRepository
	Config            *configs.Config
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{ProductRepository: deps.ProductRepository}

	router.HandleFunc("GET /products", handler.GetAll())
	router.HandleFunc("GET /product/{id}", handler.GetById())
	router.Handle("POST /product", middleware.IsAuth(handler.Create(), deps.Config))
	router.Handle("PATCH /product/{id}", middleware.IsAuth(handler.Update(), deps.Config))
	router.Handle("DELETE /product/{id}", middleware.IsAuth(handler.Delete(), deps.Config))
}

func (handler *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := handler.ProductRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, 200, products)
	}
}

func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathId := r.PathValue("id")

		id, err := strconv.ParseUint(pathId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, 200, product)
	}
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.ProductRepository.Create(
			&Product{
				Name:        body.Name,
				Description: body.Description,
				Images:      body.Images})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, 201, product)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phoneNumber, ok := r.Context().Value(middleware.CONTEXT_PHONE_NUMBER_KEY).(string)
		if ok {
			fmt.Println(phoneNumber)
		}

		pathId := r.PathValue("id")
		body, err := req.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseUint(pathId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedProduct, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, 201, updatedProduct)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathId := r.PathValue("id")

		id, err := strconv.ParseUint(pathId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = handler.ProductRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, 200, nil)
	}
}
