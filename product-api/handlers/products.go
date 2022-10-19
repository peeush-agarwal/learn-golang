// Package classification Product API
//
// Documentation for Product API
//
//		Schemes: http
//	 BasePath: /
//	 Version: 1.0.0
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/peeush-agarwal/learn-golang/product-api/data"
)

type KeyProduct struct{}

// A list of products returned in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		// Validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				w,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
