package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

import "github.com/eyedeekay/sam-forwarder"

type blah struct{}

var forwarder *samforwarder.SAMForwarder

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (b *blah) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("the echo service is responding to a request on:", forwarder.Base32())
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
	fmt.Fprintf(w, "  <p>\n")
	fmt.Fprintf(w, "    Attempting to force resource retrieval over plain https")
	fmt.Fprintf(w, "  </p>\n")
	fmt.Fprintf(w, "    <pre><code>\n")
	for key, value := range r.Header {
		log.Println(key, value)
		fmt.Fprintf(w, "Header: %s, Value: %s\n", key, value)
	}
	fmt.Fprintf(w, "    </pre></code>\n")
	fmt.Fprintf(w, "  <script type=\"application/javascript\">\n")
	fmt.Fprintf(w, "    function getIP(json) {\n")
	fmt.Fprintf(w, "      document.write(\"<pre><code>\");\n")
	fmt.Fprintf(w, "      document.write(\"My public IP address is:\", json.ip);\n")
	fmt.Fprintf(w, "      document.write(\"</pre></code>\");\n")
	fmt.Fprintf(w, "    }\n")
	fmt.Fprintf(w, "  </script>\n")
	fmt.Fprintf(w, "  <script type=\"application/javascript\" src=\"https://api.ipify.org?format=jsonp&callback=getIP\"></script>\n")
	fmt.Fprintf(w, "  </body>\n")
	fmt.Fprintf(w, "</html>\n")
	/*

	   fmt.Fprintf(w, "\n")

	*/
}

func main() {
	var err error
	rand.Seed(time.Now().UnixNano())
	log.Println("starting go echo service")
	samhost := flag.String("samhost", "sam-host", "host of the SAM to use")
	samport := flag.String("samport", "7656", "port of the SAM to use")
	host := flag.String("host", "0.0.0.0", "host to forward")
	port := flag.String("port", "9777", "port to forward")
	flag.Parse()
	if forwarder, err = samforwarder.NewSAMForwarderFromOptions(
		samforwarder.SetSaveFile(true),
		samforwarder.SetName("collude-"+randSeq(4)),
		samforwarder.SetSAMHost(*samhost),
		samforwarder.SetSAMPort(*samport),
		samforwarder.SetHost(*host),
		samforwarder.SetPort(*port),
	); err != nil {
		log.Fatal(err.Error())
	} else {
		go forwarder.Serve()
	}
	log.Println("Colluder configured on:", forwarder.Base32())
	log.Fatal(http.ListenAndServe(*host+":"+*port, &blah{}))
}
