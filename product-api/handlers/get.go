package handlers

import (
	"net/http"

	"github.com/peeush-agarwal/learn-golang/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		p.l.Println("Unable to read request body", err)

		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}
}
