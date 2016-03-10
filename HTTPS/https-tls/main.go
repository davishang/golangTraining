// create a web page which serves at localhost over https using TLS

/***************************************************************************
ListenAndServeTLS acts identically to ListenAndServe, except that it expects
HTTPS connections. Additionally, files containing a certificate and matching
private key for the server must be provided. If the certificate is signed by
a certificate authority, the certFile should be the concatenation of the
server's certificate, any intermediates, and the CA's certificate.

One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
ListenAndServeTLS always returns a non-nil error.
***************************************************************************/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, `<!DOCTYPE HTML>
		<html>
			<head>
			<title></title>
			</head>
			<body>
				Sample secure server
			</body>
		</html>
		`)
}

func redir(res http.ResponseWriter, req *http.Request){
	http.Redirect(res, req, "https://127.0.0.1:10443/" + req.RequestURI, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", foo)
	log.Println("Listening...on 10443.  Go to https://127.0.0.1:10443/")
	go http.ListenAndServe(":8080", http.HandlerFunc(redir))
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Generate unsigned certificate
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
// for example
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost

// WINDOWS
// windows may have issues with go env GOROOT
// go run %(go env GOROOT)%/src/crypto/tls/generate_cert.go --host=localhost

//************ instead of go env GOROOT
//************ you can just use the path to the GO SDK
//************ wherever it is on your computer

//find where cert.pem and key.pem are generated and saved at, copy it to folder
//containing the .go file that you're working with

// Go to https://localhost:10443/ or https://127.0.0.1:10443/

