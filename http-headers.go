package main

import (
	"fmt"
	"log"
	"net/http"
)

type blah struct{}

func (b *blah) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println("the echo service is responding.")
	for key, value := range r.Header {
		fmt.Fprintf(w, "Temos o header: %s, com valor %s\n", key, value)
	}
}

func main() {
    log.Println("starting go echo service")
	log.Fatal(http.ListenAndServe(":8080", &blah{}))
}
