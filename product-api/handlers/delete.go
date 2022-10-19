package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/peeush-agarwal/learn-golang/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete the specified product
// responses:
// 	201: noContentResponse

// DeleteProduct deletes the product from the data store
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("[ERROR] Unable to convert id into int from string", err)
		http.Error(rw, "Invalid id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle DELETE Product for", id)

	err = data.DeleteProduct(id)

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
