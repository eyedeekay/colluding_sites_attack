package main

import (
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

import "github.com/eyedeekay/sam-forwarder"

var forwarder *samforwarder.SAMForwarder

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// CSSStyle prints the contents of the CSS file
func CSSStyle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, css)
}

// FingerprintJS prints the contents of fingeprint.js
func FingerprintJS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fingerprintjs)
}

func LocalJS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `    document.querySelector("#btn").addEventListener("click", function () {%s`, "\n")
	fmt.Fprintf(w, `      var d1 = new Date();%s`, "\n")
	fmt.Fprintf(w, `      var fp = new Fingerprint2();%s`, "\n")
	fmt.Fprintf(w, `      fp.get(function(result, components) {%s`, "\n")
	fmt.Fprintf(w, `        var d2 = new Date();%s`, "\n")
	fmt.Fprintf(w, `        var timeString = "Time took to calculate the fingerprint: " + (d2 - d1) + "ms";%s`, "\n")
	fmt.Fprintf(w, `        var details = "<strong>Detailed information: </strong><br />";%s`, "\n")
	fmt.Fprintf(w, `        if(typeof window.console !== "undefined") {%s`, "\n")
	fmt.Fprintf(w, `          console.log(timeString);%s`, "\n")
	fmt.Fprintf(w, `          console.log(result);%s`, "\n")
	fmt.Fprintf(w, `          for (var index in components) {%s`, "\n")
	fmt.Fprintf(w, `            var obj = components[index];%s`, "\n")
	fmt.Fprintf(w, `            var value = obj.value;%s`, "\n")
	fmt.Fprintf(w, `            var line = obj.key + " = " + value.toString().substr(0, 100);%s`, "\n")
	fmt.Fprintf(w, `            console.log(line);%s`, "\n")
	fmt.Fprintf(w, `            details += line + "<br />";%s`, "\n")
	fmt.Fprintf(w, `          }%s`, "\n")
	fmt.Fprintf(w, `        }%s`, "\n")
	fmt.Fprintf(w, `        document.querySelector("#details").innerHTML = details%s`, "\n")
	fmt.Fprintf(w, `        document.querySelector("#fp").textContent = result%s`, "\n")
	fmt.Fprintf(w, `        document.querySelector("#time").textContent = timeString%s`, "\n")
	fmt.Fprintf(w, `      });%s`, "\n")
	fmt.Fprintf(w, `    });%s`, "\n")
}

// PageContent builds the page
func PageContent(w http.ResponseWriter, r *http.Request) {
	log.Println("the echo service is responding to a request on:", forwarder.Base32())
	csp_header := fmt.Sprintf("default-src 'self' api.ipify.org %s; ", *sourcesite)
	csp_header += fmt.Sprintf("script-src 'self' api.ipify.org %s;", *sourcesite)
	w.Header().Add("Content-Security-Policy", csp_header)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>`))
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, `<html>%s`, "\n")
	fmt.Fprintf(w, `<head>%s`, "\n")
	fmt.Fprintf(w, `  <title> What is my Base64? </title>%s`, "\n")
	if *sourcesite != "none" {
		fmt.Fprintf(w, `  <link rel="stylesheet" href="http://%s/css/styles.css">%s`, *sourcesite, "\n")
	} else {
		fmt.Fprintf(w, `  <link rel="stylesheet" href="/styles.css">%s`, "\n")
	}
	fmt.Fprintf(w, `</head>%s`, "\n")
	fmt.Fprintf(w, `  <body>%s`, "\n")
	fmt.Fprintf(w, `  <p>%s`, "\n")
	fmt.Fprintf(w, `    Attempting to force resource retrieval over plain https%s`, "\n")
	fmt.Fprintf(w, `  </p>%s`, "\n")
	fmt.Fprintf(w, `    <pre><code>%s`, "\n")
	fmt.Fprintf(w, "visited:%s\n", html.EscapeString(r.URL.Path))
	for key, value := range r.Header {
		log.Println(key, value)
		fmt.Fprintf(w, "Header: %s, Value: %s \n", key, value)
	}
	fmt.Fprintf(w, `    </pre></code>%s`, "\n")
	fmt.Fprintf(w, `  <script type="application/javascript">%s`, "\n")
	fmt.Fprintf(w, `    function getIP(json) {%s`, "\n")
	fmt.Fprintf(w, `      document.write("<pre><code>");%s`, "\n")
	fmt.Fprintf(w, `      document.write("My public IP address is:", json.ip);%s`, "\n")
	fmt.Fprintf(w, `      document.write("</pre></code>");%s`, "\n")
	fmt.Fprintf(w, `    }%s`, "\n")
	fmt.Fprintf(w, `  </script>%s`, "\n")
	fmt.Fprintf(w, `  <script type="application/javascript" src="https://api.ipify.org?format=jsonp&callback=getIP"></script>%s`, "\n")
	fmt.Fprintf(w, `  <div id="container"></div>%s`, "\n")
	fmt.Fprintf(w, `  <h3>Fingerprintjs2</h3>%s`, "\n")
	fmt.Fprintf(w, `  <p>Your browser fingerprint: <strong id="fp"></strong></p>%s`, "\n")
	fmt.Fprintf(w, `  <p><code id="time"/></p>%s`, "\n")
	fmt.Fprintf(w, `  <p><span id="details"/></p>%s`, "\n")
	fmt.Fprintf(w, `  <button type="button" id="btn">Get my fingerprint</button>%s`, "\n")
	if *sourcesite != "none" {
		fmt.Fprintf(w, `  <script type="application/javascript" src="http://%s/include/fingerprint2.js"></script>%s`, *sourcesite, "\n")
	} else {
		fmt.Fprintf(w, `  <script type="application/javascript" src="/fingerprint.js"></script>%s`, "\n")
	}
	fmt.Fprintf(w, `  <script>%s`, "\n")
	fmt.Fprintf(w, `  <script type="application/javascript" src="/local.js"></script>%s`, "\n")
	fmt.Fprintf(w, `  </script>%s`, "\n")
	fmt.Fprintf(w, `  </body>%s`, "\n")
	fmt.Fprintf(w, `</html>%s`, "\n")
}

var (
	samhost          = flag.String("samhost", "sam-host", "host of the SAM to use")
	samport          = flag.String("samport", "7656", "port of the SAM to use")
	host             = flag.String("host", "0.0.0.0", "host to forward")
	port             = flag.String("port", "9777", "port to forward")
	tag              = flag.String("tag", randSeq(4), "append to collude-* name")
	sourcesite       = flag.String("resource", "none", "b32 address of site with resources")
	toralso          = flag.Bool("tor", false, "Also deploy a Tor Onion Service and try to weaken Tor Browsing")
	fingperintjspath = flag.String("finger", "./include/fingerprint2.js", "Load fingerprintjs from this source file.")
	csspath          = flag.String("css", "./css/styles.css", "Load CSS file from this source file")
	fingerprintjs    string
	css              string
)

func main() {
	var err error
	rand.Seed(time.Now().UnixNano())
	log.Println("starting go echo service")
	flag.Parse()
	if forwarder, err = samforwarder.NewSAMForwarderFromOptions(
		samforwarder.SetSaveFile(true),
		samforwarder.SetName("collude-"+*tag),
		samforwarder.SetSAMHost(*samhost),
		samforwarder.SetSAMPort(*samport),
		samforwarder.SetHost(*host),
		samforwarder.SetPort(*port),
		samforwarder.SetType("http"),
	); err != nil {
		log.Fatal(err.Error())
	} else {
		go forwarder.Serve()
	}
	fbytes, err := ioutil.ReadFile(*fingperintjspath)
	if err != nil {
		panic(err)
	}
	fingerprintjs = string(fbytes)
	cbytes, err := ioutil.ReadFile(*csspath)
	if err != nil {
		panic(err)
	}
	css = string(cbytes)
	http.HandleFunc("/", PageContent)
	http.HandleFunc("/styles.css", CSSStyle)
	http.HandleFunc("/fingerprint.js", FingerprintJS)
	http.HandleFunc("/local.js", LocalJS)
	log.Println("Colluder configured on:", forwarder.Base32())
	log.Fatal(http.ListenAndServe(*host+":"+*port, nil))
}
