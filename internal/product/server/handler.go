package server

import (
	"net/http"

	"github.com/Ceilovee/Bootcamp-go/internal/product/servise"
)

type HandlerDefault struct {
	Sv *servise.ServiceDefault
}

func NewHandlerDefault(sv *servise.ServiceDefault) *HandlerDefault {
	return &HandlerDefault{Sv: sv}
}

func (hd *HandlerDefault) HandlerGetProducts() http.HandlerFunc {
	return hd.Sv.HandlerGetProducts()
}

func (hd *HandlerDefault) HandlerPong() http.HandlerFunc {
	return hd.Sv.HandlerPong()
}

func (hd *HandlerDefault) HandlerCreateProduct() http.HandlerFunc {
	return hd.Sv.HandlerCreateProduct()
}

func (hd *HandlerDefault) HandlerGetProductByID() http.HandlerFunc {
	return hd.Sv.HandlerGetProductByID()
}

func (hd *HandlerDefault) HandlerSearchProducts() http.HandlerFunc {
	return hd.Sv.HandlerSearchProducts()
}
