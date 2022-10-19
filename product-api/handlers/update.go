package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/peeush-agarwal/learn-golang/product-api/data"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("[ERROR] Unable to convert id into int from string", err)
		http.Error(rw, "Invalid id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product for", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] Product not found for the specified id", err)
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] Unknown error occurred while updating product", err)
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
