package main

import (
	"chat-websocket/src/cmd/auth"
	"chat-websocket/src/cmd/db"
	"encoding/json"
	"log"
	"net/http"
	"src/github.com/gorilla/mux"
	"time"
)

func main() {
	r := mux.NewRouter()
	dir := "./front/build"
	//r.Path("/static").Handler(http.FileServer(http.Dir(dir)))
	r.PathPrefix("/api/auth").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		name := request.Header.Get("name")
		pass := request.Header.Get("password")
		us, err := auth.Register(name, pass)
		if err != nil || us == nil {
			http.Error(writer, "go avay", http.StatusNotAcceptable)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		log.Println(us)
		res := auth.ResponseAuth{
			Token:   us.Token,
			Message: "OK",
		}
		data, _ := json.Marshal(res)
		_, _ = writer.Write(data)
	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer db.Close()
	log.Fatal(srv.ListenAndServe())
}
