package servise

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ceilovee/Bootcamp-go/internal/product/storage"
	"github.com/go-chi/chi"
)

type ServiceDefault struct {
	P *storage.ProductsController
}

type ResponseBodyProduct struct {
	Message string           `json:"message"`
	Data    *storage.Product `json:"data"`
	Error   string           `json:"error"`
}

type RequestBodyProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is-published,omitempty"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func NewServiceDefault(p *storage.ProductsController) *ServiceDefault {
	return &ServiceDefault{P: p}
}

func (sv *ServiceDefault) HandlerPong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

func (sv *ServiceDefault) HandlerGetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(sv.P.Prod); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

func (sv *ServiceDefault) HandlerGetProductByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var prod storage.Product
		id := chi.URLParam(r, "id")
		id_int, err := strconv.Atoi(id)
		if err != nil && id_int <= sv.P.N {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		for _, product := range sv.P.Prod {
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

func (sv *ServiceDefault) HandlerSearchProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prods := make([]storage.Product, 0)
		priceGT := r.URL.Query().Get("priceGT")
		price_float, err := strconv.ParseFloat(priceGT, 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		for _, product := range sv.P.Prod {
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

func (sv *ServiceDefault) HandlerCreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// token := w.Header().Get("token")
		// fmt.Println(token)

		// if token != "marte" {
		// 	err := fmt.Errorf("Invalid token, access denied")
		// 	fmt.Println(err)
		// 	errorMessage(w, err)
		// 	return
		// }

		var rb RequestBodyProduct
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&rb); err != nil {
			errorMessage(w, err)
			return
		}

		pp, err := sv.P.AddProduct(rb.Name, rb.Expiration, rb.CodeValue, rb.IsPublished, rb.Price, rb.Quantity)

		if err != nil {
			errorMessage(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		body := ResponseBodyProduct{
			Message: "Product created successfully",
			Data:    &pp,
			Error:   "",
		}
		json.NewEncoder(w).Encode(body)
	}
}

func errorMessage(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	body := ResponseBodyProduct{
		Message: "Bad request",
		Data:    nil,
		Error:   err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
	return
}
