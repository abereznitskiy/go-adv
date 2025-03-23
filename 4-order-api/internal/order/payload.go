package order

import "go-adv/4-order-api/pkg/models"

type OrderCreateRequest struct {
	Products []*models.Product `json:"products" validate:"required"`
}
