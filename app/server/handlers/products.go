package handlers

import (
	"challenge/internal/product"
	"challenge/utils"
	"fmt"
	"net/http"
)

type ProductHandler struct {
	service *product.Service
}

func NewProductHandler() *ProductHandler {
	service := product.NewService()

	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.service.GetProducts()

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve products", fmt.Errorf("failed to retrieve products: %v", err))
			return
		}

		utils.ReturnJson(w, products)
	}
}
