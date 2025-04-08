package internal

import (
	"net/http"

	"github.com/go-chi/chi"
)

func RaiseServer(p *ProductsController) {
	r := chi.NewRouter()

	r.Get("/ping", handlerPong)
	r.Get("/products", p.handlerGetProducts())
	r.Post("/products", p.handlerCreateProduct())
	r.Get("/products/{id}", p.handlerGetProductByID())
	r.Get("/products/search", p.handlerSearchProducts())

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
