package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/peeush-agarwal/learn-golang/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Products requests")

	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		p.l.Println("Unable to read request body", err)

		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}
