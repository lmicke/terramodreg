package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lmicke/terramodreg/registry"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/.well-known/terraform.json", registry.ServiceDiscoveryHandler)

	//Module Handler
	/*
		s := r.PathPrefix("/terraform/modules/v1").SubRouter()
		s.Methods("GET")
		s.HandleFunc("/{namespace}/{provider}/versions", VersionHandler)
		s.HandleFunc("/{namespace}/{provider}/{version}/download", DownloadHandler)
	*/

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
