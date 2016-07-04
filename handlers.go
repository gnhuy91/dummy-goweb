package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Welcome!")
	j, _ := json.Marshal(&Message{s})
	w.Write(j)
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	m := Message{"go API, build v0.0.001.992."}
	j, _ := json.Marshal(&m)
	w.Write(j)
}
