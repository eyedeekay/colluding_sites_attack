package echosam

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (f *EchoSAM) FindNext(filename string) string {
	num := 0
	split := strings.SplitN(filename, ".", 2)
	if _, err := ioutil.ReadFile(filename); err != nil {
		return filename
	}
	for {
		if _, err := ioutil.ReadFile(split[0] + "-" + strconv.Itoa(num) + "." + split[1]); err != nil {
			return split[0] + "-" + strconv.Itoa(num) + "." + split[1]
		}
	}
}

// Report logs information about an experimental participant after they
// voluntarily opt-in. Right now it does nothing.
func (f *EchoSAM) Report(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Fprintf(w, string(body))
		fmt.Fprintf(w, "\n")
		err := ioutil.WriteFile(f.FindNext("file.txt"), body, 0644)
		if err != nil {
			log.Println("ERROR: ", err.Error())
		}
	} else {
		log.Println("ERROR: ", err.Error())
		fmt.Fprintf(w, err.Error())
		fmt.Fprintf(w, "\n")
		err := ioutil.WriteFile(f.FindNext("file.txt"), []byte(err.Error()), 0644)
		if err != nil {
			log.Println("ERROR: ", err.Error())
		}
	}
}

// CSSStyle prints the contents of the CSS file
func (f *EchoSAM) CSSStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	cbytes, _ := ioutil.ReadFile(f.CSS)
	css := string(cbytes)
	fmt.Fprintf(w, css)
	fmt.Fprintf(w, "\n")
}

// FingerprintJS prints the contents of fingeprint.js
func (f *EchoSAM) Fingerprint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	fbytes, _ := ioutil.ReadFile(f.FingerprintJS)
	fingerprintjs := string(fbytes)
	fmt.Fprintf(w, fingerprintjs)
	fmt.Fprintf(w, "\n")
}

// FingerprintJS prints the contents of fingeprint.js
func (f *EchoSAM) Finger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	f2bytes, _ := ioutil.ReadFile(f.FingerFile)
	fingerfile := string(f2bytes)
	fmt.Fprintf(w, fingerfile)
	fmt.Fprintf(w, "\n")
}

// GetIP launches the browser misconfiguration detecting script
func (f *EchoSAM) GetIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprintf(w, `    function getIP(json) {%s`, "\n")
	fmt.Fprintf(w, `      document.write("<pre><code>");%s`, "\n")
	fmt.Fprintf(w, `      document.write("My public IP address is:", json.ip);%s`, "\n")
	fmt.Fprintf(w, `      document.write("</pre></code>");%s`, "\n")
	fmt.Fprintf(w, `    }%s`, "\n")
}

// LocalJS loads the on-page components of fingerprintjs
func (f *EchoSAM) Local(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	lbytes, _ := ioutil.ReadFile(f.LocalJS)
	//if err != nil {
	//panic(err)
	//}
	localjs := string(lbytes)
	fmt.Fprintf(w, localjs)
	fmt.Fprintf(w, "\n")
}

// FingerSection prints the FingerprintJS section of the page.
func (f *EchoSAM) FingerSection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `  <div id="fingerprintjs">%s`, "\n")
	fmt.Fprintf(w, `  <a href="/finger.html">Get fingerprint</a>%s`, "\n")
	fmt.Fprintf(w, `  </div>%s`, "\n")
}

// IPSection prints the browser misconfiguration IP leak section
func (f *EchoSAM) IPSection(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, `    <script src="/getip.js"></script>%s`, "\n")
	fmt.Fprintf(w, `  </div>%s`, "\n")
}

func (f *EchoSAM) HeaderSection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>`))
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, `<html>%s`, "\n")
	fmt.Fprintf(w, `<head>%s`, "\n")
	fmt.Fprintf(w, `  <title> What is my Base64? </title>%s`, "\n")
	fmt.Fprintf(w, `  <link rel="stylesheet" type="text/css" href="/styles.css">%s`, "\n")
	fmt.Fprintf(w, `</head>%s`, "\n")
}

// PageContent builds the page
func (f *EchoSAM) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("the echo service is responding to a request on:", f.Base32())
	if strings.HasSuffix(r.URL.Path, "styles.css") {
		f.CSSStyle(w, r)
		return
	}
	if strings.HasSuffix(r.URL.Path, "client.js") || strings.HasSuffix(r.URL.Path, "client.min.js") {
		f.Fingerprint(w, r)
		return
	}
	if strings.HasSuffix(r.URL.Path, "local.js") {
		f.Local(w, r)
		return
	}
	if strings.HasSuffix(r.URL.Path, "getip.js") {
		f.GetIP(w, r)
		return
	}
	if strings.HasSuffix(r.URL.Path, "finger.html") {
		f.Finger(w, r)
		return
	}
	if strings.HasSuffix(r.URL.Path, "report") {
		f.Report(w, r)
		return
	}
	f.HeaderSection(w, r)
	fmt.Fprintf(w, `  <body>%s`, "\n")
	f.IPSection(w, r)
	f.FingerSection(w, r)
	fmt.Fprintf(w, `  <script>%s`, "\n")
	fmt.Fprintf(w, `  </script>%s`, "\n")
	fmt.Fprintf(w, `  </body>%s`, "\n")
	fmt.Fprintf(w, `</html>%s`, "\n")

}
