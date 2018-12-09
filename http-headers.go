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

// GetIP launches the browser misconfiguration detecting script
func GetIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `    function getIP(json) {%s`, "\n")
	fmt.Fprintf(w, `      document.write("<pre><code>");%s`, "\n")
	fmt.Fprintf(w, `      document.write("My public IP address is:", json.ip);%s`, "\n")
	fmt.Fprintf(w, `      document.write("</pre></code>");%s`, "\n")
	fmt.Fprintf(w, `    }%s`, "\n")
}

// LocalJS loads the on-page components of fingerprintjs
func LocalJS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, localjs)
}

// FingerSection prints the FingerprintJS section of the page.
func FingerSection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `  <div id="fingerprintjs">%s`, "\n")
	fmt.Fprintf(w, `    <h3>Fingerprintjs2</h3>%s`, "\n")
	fmt.Fprintf(w, `    <p>Your browser fingerprint: <strong id="fp"></strong></p>%s`, "\n")
	fmt.Fprintf(w, `    <p><code id="time"/></p>%s`, "\n")
	fmt.Fprintf(w, `    <p><span id="details"/></p>%s`, "\n")
	fmt.Fprintf(w, `    <button type="button" id="btn">Get my fingerprint</button>%s`, "\n")
	if *sourcesite != "none" {
		fmt.Fprintf(w, `    <script type="application/javascript" src="http://%s/include/fingerprint2.js"></script>%s`, *sourcesite, "\n")
	} else {
		fmt.Fprintf(w, `    <script type="application/javascript" src="/fingerprint.js"></script>%s`, "\n")
	}
	fmt.Fprintf(w, `    <script defer type="application/javascript" src="/local.js"></script>%s`, "\n")
	fmt.Fprintf(w, `  </div>%s`, "\n")
}

// IPSection prints the browser misconfiguration IP leak section
func IPSection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `  <div id="browsertest">%s`, "\n")
	fmt.Fprintf(w, `    <p>%s`, "\n")
	fmt.Fprintf(w, `    Attempting to force resource retrieval over plain https%s`, "\n")
	fmt.Fprintf(w, `    </p>%s`, "\n")
	fmt.Fprintf(w, `      <pre><code>%s`, "\n")
	fmt.Fprintf(w, "  visited:%s\n", html.EscapeString(r.URL.Path))
	for key, value := range r.Header {
		log.Println(key, value)
		fmt.Fprintf(w, "Header: %s, Value: %s \n", key, value)
	}
	fmt.Fprintf(w, `      </pre></code>%s`, "\n")
	fmt.Fprintf(w, `    <script type="application/javascript" src="/getip.js"></script>%s`, "\n")
	fmt.Fprintf(w, `    <script type="application/javascript" src="https://api.ipify.org?format=jsonp&callback=getIP"></script>%s`, "\n")
	fmt.Fprintf(w, `  </div>%s`, "\n")
}

func HeaderSection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>`))
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, `<html>%s`, "\n")
	fmt.Fprintf(w, `<head>%s`, "\n")
	fmt.Fprintf(w, `  <title> What is my Base64? </title>%s`, "\n")
	if *sourcesite != "none" {
		fmt.Fprintf(w, `  <link rel="stylesheet" href="http://%s/css/styles.css">%s`, *sourcesite, "\n")
	} else {
		fmt.Fprintf(w, `  <link rel="stylesheet" type="text/css" href="/styles.css">%s`, "\n")
	}
	fmt.Fprintf(w, `</head>%s`, "\n")
}

// PageContent builds the page
func PageContent(w http.ResponseWriter, r *http.Request) {
	log.Println("the echo service is responding to a request on:", forwarder.Base32())
	HeaderSection(w, r)
	fmt.Fprintf(w, `  <body>%s`, "\n")
	IPSection(w, r)
	FingerSection(w, r)
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
	jspath           = flag.String("js", "./include/local.js", "Load local javascript from this source file.")
	csspath          = flag.String("css", "./css/styles.css", "Load CSS file from this source file")
	fingerprintjs    string
	localjs          string
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
	lbytes, err := ioutil.ReadFile(*jspath)
	if err != nil {
		panic(err)
	}
	localjs = string(lbytes)
	http.HandleFunc("/", PageContent)
	http.HandleFunc("/styles.css", CSSStyle)
	http.HandleFunc("/fingerprint.js", FingerprintJS)
	http.HandleFunc("/local.js", LocalJS)
	http.HandleFunc("/getip.js", GetIP)
	log.Println("Colluder configured on:", forwarder.Base32())
	log.Fatal(http.ListenAndServe(*host+":"+*port, nil))
}
