package main

import (
	"testing"

	"github.com/peeush-agarwal/learn-golang/product-api-client/client"
	"github.com/peeush-agarwal/learn-golang/product-api-client/client/products"
)

func TestProductsClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	listProducts, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", listProducts.GetPayload()[0])
}
