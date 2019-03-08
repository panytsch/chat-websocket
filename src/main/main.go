package main

import (
	"chat-websocket/src/db"
	"log"
	"net/http"
	"src/github.com/gorilla/mux"
	"time"
)

func main() {
	d := db.GetDB()
	defer d.Close()

	r := mux.NewRouter()
	dir := "./front/build"
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
