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
	fmt.Fprintf(w, `<!DOCTYPE html>`)
	fmt.Fprintf(w, `<html>`)
	fmt.Fprintf(w, `<head>`)
	fmt.Fprintf(w, `  <title> What is my Base64? </title>`)
	fmt.Fprintf(w, `  <link rel="stylesheet" href="http://zbprsnu26qqtm3cccx6imtm4rc2v3o474eezdd2mfj7x4bpcnqqq.b32.i2p/css/styles.css">`)
	fmt.Fprintf(w, `</head>`)
	fmt.Fprintf(w, `  <body>`)
	fmt.Fprintf(w, `  <p>`)
	fmt.Fprintf(w, `    Attempting to force resource retrieval over plain https`)
	fmt.Fprintf(w, `  </p>`)
	fmt.Fprintf(w, `    <pre><code>`)
	for key, value := range r.Header {
		log.Println(key, value)
		fmt.Fprintf(w, "Header: %s, Value: %sn", key, value)
	}
	fmt.Fprintf(w, `    </pre></code>`)
	fmt.Fprintf(w, `  <script type="application/javascript">`)
	fmt.Fprintf(w, `    function getIP(json) {`)
	fmt.Fprintf(w, `      document.write("<pre><code>");`)
	fmt.Fprintf(w, `      document.write("My public IP address is:", json.ip);`)
	fmt.Fprintf(w, `      document.write("</pre></code>");`)
	fmt.Fprintf(w, `    }`)
	fmt.Fprintf(w, `  </script>`)
	fmt.Fprintf(w, `  <script type="application/javascript" src="https://api.ipify.org?format=jsonp&callback=getIP"></script>`)
    fmt.Fprintf(w, `  <div id="container"></div>`)
    fmt.Fprintf(w, `  <h3>Fingerprintjs2</h3>`)
    fmt.Fprintf(w, `  <p>Your browser fingerprint: <strong id="fp"></strong></p>`)
    fmt.Fprintf(w, `  <p><code id="time"/></p>`)
    fmt.Fprintf(w, `  <p><span id="details"/></p>`)
    fmt.Fprintf(w, `  <button type="button" id="btn">Get my fingerprint</button>`)

    fmt.Fprintf(w, `  <script src="http://zbprsnu26qqtm3cccx6imtm4rc2v3o474eezdd2mfj7x4bpcnqqq.b32.i2p/include/fingerprint2.js"></script>`)
    fmt.Fprintf(w, `  <script>`)
    fmt.Fprintf(w, `    document.querySelector("#btn").addEventListener("click", function () {`)
    fmt.Fprintf(w, `      var d1 = new Date();`)
    fmt.Fprintf(w, `      var fp = new Fingerprint2();`)
    fmt.Fprintf(w, `      fp.get(function(result, components) {`)
    fmt.Fprintf(w, `        var d2 = new Date();`)
    fmt.Fprintf(w, `        var timeString = "Time took to calculate the fingerprint: " + (d2 - d1) + "ms";`)
    fmt.Fprintf(w, `        var details = "<strong>Detailed information: </strong><br />";`)
    fmt.Fprintf(w, `        if(typeof window.console !== "undefined") {`)
    fmt.Fprintf(w, `          console.log(timeString);`)
    fmt.Fprintf(w, `          console.log(result);`)
    fmt.Fprintf(w, `          for (var index in components) {`)
    fmt.Fprintf(w, `            var obj = components[index];`)
    fmt.Fprintf(w, `            var value = obj.value;`)
    fmt.Fprintf(w, `            var line = obj.key + " = " + value.toString().substr(0, 100);`)
    fmt.Fprintf(w, `            console.log(line);`)
    fmt.Fprintf(w, `            details += line + "<br />";`)
    fmt.Fprintf(w, `          }`)
    fmt.Fprintf(w, `        }`)
    fmt.Fprintf(w, `        document.querySelector("#details").innerHTML = details`)
    fmt.Fprintf(w, `        document.querySelector("#fp").textContent = result`)
    fmt.Fprintf(w, `        document.querySelector("#time").textContent = timeString`)
    fmt.Fprintf(w, `      });`)
    fmt.Fprintf(w, `    });`)
    fmt.Fprintf(w, `  </script>`)
	fmt.Fprintf(w, `  </body>`)
	fmt.Fprintf(w, `</html>`)
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
