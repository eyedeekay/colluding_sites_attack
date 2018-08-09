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
	fmt.Fprintf(w, `<!DOCTYPE html>`, "\n")
	fmt.Fprintf(w, `<html>`, "\n")
	fmt.Fprintf(w, `<head>`, "\n")
	fmt.Fprintf(w, `  <title> What is my Base64? </title>`, "\n")
	fmt.Fprintf(w, `  <link rel="stylesheet" href="http://zbprsnu26qqtm3cccx6imtm4rc2v3o474eezdd2mfj7x4bpcnqqq.b32.i2p/css/styles.css">`, "\n")
	fmt.Fprintf(w, `</head>`, "\n")
	fmt.Fprintf(w, `  <body>`, "\n")
	fmt.Fprintf(w, `  <p>`, "\n")
	fmt.Fprintf(w, `    Attempting to force resource retrieval over plain https`, "\n")
	fmt.Fprintf(w, `  </p>`, "\n")
	fmt.Fprintf(w, `    <pre><code>`, "\n")
	for key, value := range r.Header {
		log.Println(key, value)
		fmt.Fprintf(w, "Header: %s, Value: %s \n", key, value)
	}
	fmt.Fprintf(w, `    </pre></code>`, "\n")
	fmt.Fprintf(w, `  <script type="application/javascript">`, "\n")
	fmt.Fprintf(w, `    function getIP(json) {`, "\n")
	fmt.Fprintf(w, `      document.write("<pre><code>");`, "\n")
	fmt.Fprintf(w, `      document.write("My public IP address is:", json.ip);`, "\n")
	fmt.Fprintf(w, `      document.write("</pre></code>");`, "\n")
	fmt.Fprintf(w, `    }`, "\n")
	fmt.Fprintf(w, `  </script>`, "\n")
	fmt.Fprintf(w, `  <script type="application/javascript" src="https://api.ipify.org?format=jsonp&callback=getIP"></script>`, "\n")
    fmt.Fprintf(w, `  <div id="container"></div>`, "\n")
    fmt.Fprintf(w, `  <h3>Fingerprintjs2</h3>`, "\n")
    fmt.Fprintf(w, `  <p>Your browser fingerprint: <strong id="fp"></strong></p>`, "\n")
    fmt.Fprintf(w, `  <p><code id="time"/></p>`, "\n")
    fmt.Fprintf(w, `  <p><span id="details"/></p>`, "\n")
    fmt.Fprintf(w, `  <button type="button" id="btn">Get my fingerprint</button>`, "\n")

    fmt.Fprintf(w, `  <script src="http://zbprsnu26qqtm3cccx6imtm4rc2v3o474eezdd2mfj7x4bpcnqqq.b32.i2p/include/fingerprint2.js"></script>`, "\n")
    fmt.Fprintf(w, `  <script>`, "\n")
    fmt.Fprintf(w, `    document.querySelector("#btn").addEventListener("click", function () {`, "\n")
    fmt.Fprintf(w, `      var d1 = new Date();`, "\n")
    fmt.Fprintf(w, `      var fp = new Fingerprint2();`, "\n")
    fmt.Fprintf(w, `      fp.get(function(result, components) {`, "\n")
    fmt.Fprintf(w, `        var d2 = new Date();`, "\n")
    fmt.Fprintf(w, `        var timeString = "Time took to calculate the fingerprint: " + (d2 - d1) + "ms";`, "\n")
    fmt.Fprintf(w, `        var details = "<strong>Detailed information: </strong><br />";`, "\n")
    fmt.Fprintf(w, `        if(typeof window.console !== "undefined") {`, "\n")
    fmt.Fprintf(w, `          console.log(timeString);`, "\n")
    fmt.Fprintf(w, `          console.log(result);`, "\n")
    fmt.Fprintf(w, `          for (var index in components) {`, "\n")
    fmt.Fprintf(w, `            var obj = components[index];`, "\n")
    fmt.Fprintf(w, `            var value = obj.value;`, "\n")
    fmt.Fprintf(w, `            var line = obj.key + " = " + value.toString().substr(0, 100);`, "\n")
    fmt.Fprintf(w, `            console.log(line);`, "\n")
    fmt.Fprintf(w, `            details += line + "<br />";`, "\n")
    fmt.Fprintf(w, `          }`, "\n")
    fmt.Fprintf(w, `        }`, "\n")
    fmt.Fprintf(w, `        document.querySelector("#details").innerHTML = details`, "\n")
    fmt.Fprintf(w, `        document.querySelector("#fp").textContent = result`, "\n")
    fmt.Fprintf(w, `        document.querySelector("#time").textContent = timeString`, "\n")
    fmt.Fprintf(w, `      });`, "\n")
    fmt.Fprintf(w, `    });`, "\n")
    fmt.Fprintf(w, `  </script>`, "\n")
	fmt.Fprintf(w, `  </body>`, "\n")
	fmt.Fprintf(w, `</html>`, "\n")
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
