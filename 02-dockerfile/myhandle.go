package main

import (
	"log"
	"net/http"
	"os"
)

type MyHandler struct {
	version string
}

func NewMyHandler() *MyHandler {
	version, _ := os.LookupEnv("VERSION")
	return &MyHandler{version: version}
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	switch r.URL.Path {
	case "/healthz":
		h.healthzFunc(w, r)
	case "/version":
		code = h.versionFunc(w, r)
	default:
		code = defaultFunc(w, r)
	}
	log.Printf("%s\t%s\t%d", r.URL.Path, r.RemoteAddr, code)

}

func (h *MyHandler) healthzFunc(w http.ResponseWriter, r *http.Request) int {
	w.WriteHeader(http.StatusOK)
	return http.StatusOK
}

func (h *MyHandler) versionFunc(w http.ResponseWriter, r *http.Request) int {
	resultCode := http.StatusOK
	if h.version == "" {
		resultCode = http.StatusInternalServerError
		w.WriteHeader(resultCode)
		w.Write([]byte("Unknown version!"))
	} else {
		w.Header().Set("version", h.version)
		w.Write([]byte(h.version))
	}
	return resultCode
}

func defaultFunc(resp http.ResponseWriter, req *http.Request) int {
	for k, vs := range req.Header {
		for _, v := range vs {
			resp.Header().Add(k, v)
		}
	}
	return http.StatusOK
}
