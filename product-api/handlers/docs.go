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

import "github.com/peeush-agarwal/learn-golang/product-api/data"

// A list of products returned in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// A single product returned in the response
// swagger:response productResponse
type productResponseWrapper struct {
	// Single product in the system
	// in: body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters deleteProduct updateProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters addProduct
type productParamsWrapper struct {
	// Product data structure to create or update
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}
