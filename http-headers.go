package main

import (
	"fmt"
	"log"
	"net/http"
)

type blah struct{}

func (b *blah) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Fprintf(w, "Temos o header: %s, com valor %s\n", key, value)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &blah{}))
}
