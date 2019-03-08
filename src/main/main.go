package main

import (
	"chat-websocket/src/auth"
	"chat-websocket/src/db"
	"encoding/json"
	"log"
	"net/http"
	"src/github.com/gorilla/mux"
	"time"
)

func main() {
	d := db.GetDB()
	defer d.Close()

	r := mux.NewRouter()
	//dir := "./front/build"
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
	r.PathPrefix("/api/auth").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		name := request.Header.Get("name")
		pass := request.Header.Get("password")
		us, err := auth.Register(name, pass)
		writer.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(writer, "go avay", http.StatusNotAcceptable)
		}
		res := auth.ResponseAuth{
			Token:   us.Token,
			Message: "OK",
		}
		data, _ := json.Marshal(res)
		_, _ = writer.Write(data)
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
