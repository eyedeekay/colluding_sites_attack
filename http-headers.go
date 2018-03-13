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

    */
    fmt.Fprintf(w, "<!DOCTYPE html>\n")
    fmt.Fprintf(w, "<html>\n")
    fmt.Fprintf(w, "<head>\n")
    fmt.Fprintf(w, "  <title> What is my Base64? </title>\n")
    fmt.Fprintf(w, "  <link rel=\"stylesheet\" href=\"http://4pvyyb3phqznc6e6fjewty2fpbb4p3ub2q27ojheitcg4nai6p5q.b32.i2p/css/styles.css\">\n")
    fmt.Fprintf(w, "</head>\n")

    fmt.Fprintf(w, "  <body>\n")
    fmt.Fprintf(w, "  <p>Attempting to force resource retrieval over plain https</p>\n")
    fmt.Fprintf(w, "  <iframe src=\"https://api.ipify.org\" />\n")
    fmt.Fprintf(w, "  <script type=\"application/javascript\">\n")
    fmt.Fprintf(w, "    function getIP(json) {\n")
    fmt.Fprintf(w, "      document.write(\"My public IP address is: \", json.ip);\n")
    fmt.Fprintf(w, "    }\n")
    fmt.Fprintf(w, "  </script>\n")

    fmt.Fprintf(w, "  <script type=\"application/javascript\" src=\"https://api.ipify.org?format=jsonp&callback=getIP\"></script>")
    fmt.Fprintf(w, "    <pre><code>\n")
	for key, value := range r.Header {
        log.Println(key, value)
		fmt.Fprintf(w, "Header: %s, Value: %s\n", key, value)
	}
    fmt.Fprintf(w, "    </pre></code>\n")
    fmt.Fprintf(w, "  </body>\n")
    fmt.Fprintf(w, "</html>\n")
    /*

    fmt.Fprintf(w, "\n")

    */
}

func main() {
    log.Println("starting go echo service")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", &blah{}))
}
