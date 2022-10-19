package handlers

import (
	"net/http"

	"github.com/peeush-agarwal/learn-golang/product-api/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
}
