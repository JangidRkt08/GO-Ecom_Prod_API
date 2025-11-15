package products

import (
	"log"
	"net/http"

	"github.com/jangidRkt08/go-Ecom_Prod-API/internal/json"
)


type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. CALL SERVICE -> List Products
	// 2. Return JSON in HTTP response
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// products := struct{
	// 	Products []string `json:"products"`
	// }{}

	json.Write(w, http.StatusOK,products)
	


}

