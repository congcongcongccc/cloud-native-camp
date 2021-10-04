package main

import (
	"net/http"
)

func main() {
	myHandler := NewMyHandler()
	mux := http.NewServeMux()
	mux.Handle("/", myHandler)
	http.ListenAndServe(":80", mux)
}