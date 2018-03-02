package main

import (
	"fmt"
	"log"
	"net/http"
)

type blah struct{}

func (b *blah) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println("the echo service is responding.")
    /*

    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "\n")

    */
	for key, value := range r.Header {
        log.Println(key, value)
		fmt.Fprintf(w, "Header: %s, Value: %s\n", key, value)
	}

    /*

    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "\n")

    */
}

func main() {
    log.Println("starting go echo service")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", &blah{}))
}
