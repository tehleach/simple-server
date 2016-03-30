package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/tehleach/gorelic"
)

func main() {
	agent := gorelic.NewAgent()
	agent.Verbose = true
	agent.NewrelicLicense = "e1c2bcfa464edc57cb2791a11b440ee18b05a96b"
	agent.NewrelicName = "kleach - simple server"
	agent.CollectHTTPStat = true

	http.HandleFunc("/foo", agent.WrapHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}, "foo"))
	http.HandleFunc("/bar", agent.WrapHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Goodbye, %q", html.EscapeString(r.URL.Path))
	}, "bar"))
	http.HandleFunc("/", agent.WrapHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		fmt.Fprint(w, "Welcome!")
	}, "index"))
	agent.Run()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
