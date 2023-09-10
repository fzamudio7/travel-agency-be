package main

import (
	"errors"
	"fmt"
	"net/http"
)

const serverPort = 3333

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Printf("server: %s /\n", r.Method)
			fmt.Fprintf(w, `{"airline": "delta"}`)
		}

	})
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}

}
