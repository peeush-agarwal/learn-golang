package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle GoodBye requests")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil{
		g.l.Println("Error reading request body", err)

		http.Error(rw, "Unable to read request body.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Goodbye %s\n", d)
}
