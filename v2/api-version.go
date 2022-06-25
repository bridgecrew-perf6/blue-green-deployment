package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type NewVersion struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Describe string `json:"describe"`
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	version := NewVersion{Id: "2.0.0", Name: "Version Two", Describe: "Describe New Version"}
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
