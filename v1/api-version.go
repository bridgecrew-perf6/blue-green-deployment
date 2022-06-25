package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Version struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	version := Version{Id: "1.0.0", Name: "Version One"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}

func main() {
	path := "/api/version"
	r := mux.NewRouter()

	println("app v2 is running !!!")

	r.HandleFunc(path, getVersion).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
