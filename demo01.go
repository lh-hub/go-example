package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func sayHelloGo(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			//println(name, "-------", h)
			//fmt.Fprintf(w, "%v: %v\n", name, h)
			w.Header().Set(name, h)
		}
	}
	var VERSION string
	VERSION = os.Getenv("VERSION")
	w.Header().Set("VERSION", VERSION)
	fmt.Fprintf(w, "Hello Go!")
}

func sayHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
}

func main() {
	http.HandleFunc("/", sayHelloGo)
	http.HandleFunc("/healthz", sayHealth)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
