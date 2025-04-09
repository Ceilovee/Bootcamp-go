package main

import (
	"net/http"

	"github.com/Ceilovee/Bootcamp-go/cmd/server/handler"
	"github.com/Ceilovee/Bootcamp-go/internal/product/servise"
	"github.com/Ceilovee/Bootcamp-go/internal/product/storage"
	"github.com/go-chi/chi"
)

func main() {
	st := storage.StorageInMemory()
	sv := servise.NewServiceDefault(st)
	hd := handler.NewHandlerDefault(sv)
	// server and run
	raiseServer(hd)
}

func raiseServer(hd *handler.HandlerDefault) {
	r := chi.NewRouter()

	r.Get("/ping", hd.HandlerPong())
	r.Route("/products", func(f chi.Router) {
		r.Get("", hd.HandlerGetProducts())
		r.Post("", hd.HandlerCreateProduct())
	})
	r.Get("/products/{id}", hd.HandlerGetProductByID())
	r.Get("/products/search", hd.HandlerSearchProducts())

	// run
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
