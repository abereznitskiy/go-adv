package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string         `json:"name" validate:"required,min=3,max=50"`
	Description string         `json:"description" validate:"min=5,max=100"`
	Images      pq.StringArray `json:"images" customValidate:"strings_array"`
}

type ProductUpdateRequest struct {
	Name        string         `json:"name" validate:"required,min=3,max=50"`
	Description string         `json:"description" validate:"min=5,max=100"`
	Images      pq.StringArray `json:"images" customValidate:"strings_array"`
}
