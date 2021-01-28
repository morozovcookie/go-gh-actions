package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		_, _ = io.Copy(w, bytes.NewBufferString("pong"))
	})

	_ = http.ListenAndServe("0.0.0.0:8080", mux)
}
