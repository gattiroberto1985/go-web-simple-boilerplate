package main

import (
	"log"
	"net/http"
	// app packages
	"it/bob/go/boilerplate/webapi"
)

func main() {
	http.HandleFunc("/api/sayhello", webapi.SimpleBeanHandler)
	//http.HandleFunc("/api/repo/", webapi.Repo)
	log.Fatal(http.ListenAndServe("localhost:8989", nil))
}
