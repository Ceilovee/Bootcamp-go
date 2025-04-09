package handler

import (
	"net/http"

	"github.com/Ceilovee/Bootcamp-go/internal/product/servise"
)

type HandlerDefault struct {
	sv *servise.ServiceDefault
}

func NewHandlerDefault(sv *servise.ServiceDefault) *HandlerDefault {
	return &HandlerDefault{}
}

func (hd *HandlerDefault) HandlerGetProducts() http.HandlerFunc {
	return hd.sv.HandlerGetProducts()
}

func (hd *HandlerDefault) HandlerPong() http.HandlerFunc {
	return hd.sv.HandlerPong()
}

func (hd *HandlerDefault) HandlerCreateProduct() http.HandlerFunc {
	return hd.sv.HandlerCreateProduct()
}

func (hd *HandlerDefault) HandlerGetProductByID() http.HandlerFunc {
	return hd.sv.HandlerGetProductByID()
}

func (hd *HandlerDefault) HandlerSearchProducts() http.HandlerFunc {
	return hd.sv.HandlerSearchProducts()
}
