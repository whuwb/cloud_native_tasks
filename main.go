package main

import (
	"CloudNative/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/header", getHeader)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func getHeader(w http.ResponseWriter, r *http.Request) {
	utils.UpdateHeader(w.Header(), r.Header)

	version := utils.GetVersion("default")
	w.Header().Add("VERSION", version)

	_, err := io.WriteString(w, "header updated with version: "+version)

	if err != nil {
		fmt.Printf("[Error]io write: %v", err)
	}

	trace("header", r)
}

func trace(url string, r *http.Request) {
	fmt.Printf("[tracing] %s visit %s at %s", r.RemoteAddr, url, time.Now())
}
