package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Message struct {
	Text string `json:"msg"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/about", about)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
