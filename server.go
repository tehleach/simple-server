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
	agent.CollectHTTPErrors = true

	agent.RegisterHTTPPath("foo")
	http.HandleFunc("/foo", agent.WrapHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		agent.HTTPPathErrorCounters["foo"][200].Inc(1)
		agent.HTTPErrorCounters[200].Inc(1)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}))
	agent.RegisterHTTPPath("bar")
	http.HandleFunc("/bar", agent.WrapHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		agent.HTTPErrorCounters[201].Inc(1)
		agent.HTTPPathErrorCounters["bar"][201].Inc(1)
		fmt.Fprintf(w, "Goodbye, %q", html.EscapeString(r.URL.Path))
	}))
	agent.RegisterHTTPPath("index")
	http.HandleFunc("/", agent.WrapHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		agent.HTTPErrorCounters[400].Inc(1)
		agent.HTTPPathErrorCounters["index"][400].Inc(1)
		fmt.Fprint(w, "Welcome!")
	}))
	agent.Run()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
