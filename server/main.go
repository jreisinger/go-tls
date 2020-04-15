package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jreisinger/go-tls/utils"
)

func main() {
	server := getServer()
	http.HandleFunc("/", myHandler)
	must(server.ListenAndServeTLS("", ""))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request from %s ...", r.RemoteAddr)
	w.Write([]byte("Hello from server!\n"))
	log.Printf("... body sent to %s.", r.RemoteAddr)
}

func getServer() *http.Server {
	tls := &tls.Config{
		GetCertificate: utils.CertReqFunc("cacert.crt", "cakey.key"),
	}
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tls,
	}
	return server
}

// --- helper functions ---

func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v", err)
		os.Exit(1)
	}
}
