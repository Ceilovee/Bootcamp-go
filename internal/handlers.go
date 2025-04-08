package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ResponseBodyProduct struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
	Error   error    `json:"error"`
}

func handlerPong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (p *ProductsController) handlerGetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(p.Prod); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

func (p *ProductsController) handlerGetProductByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var prod Product
		id := chi.URLParam(r, "id")
		id_int, err := strconv.Atoi(id)
		if err != nil && id_int <= p.n {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		for _, product := range p.Prod {
			if product.ID == id_int {
				prod = product
				break // asumo que solo existe un producto con ese id
			}
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(prod); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

func (p *ProductsController) handlerSearchProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prods := make([]Product, 0)
		priceGT := r.URL.Query().Get("priceGT")
		price_float, err := strconv.ParseFloat(priceGT, 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		for _, product := range p.Prod {
			if product.Price > price_float {
				prods = append(prods, product)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(prods); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

func (p *ProductsController) handlerCreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pp Product
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&pp); err != nil {
			ErrorMessage(w, err)
			return
		}
		err := p.AddProduct(pp)
		if err != nil {
			ErrorMessage(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		body := ResponseBodyProduct{
			Message: "Product created successfully",
			Data:    &pp,
			Error:   nil,
		}
		json.NewEncoder(w).Encode(body)
	}
}

func ErrorMessage(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	body := ResponseBodyProduct{
		Message: "Bad request",
		Data:    nil,
		Error:   err,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
	return
}
