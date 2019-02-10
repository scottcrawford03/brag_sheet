package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type AllBragsResponse struct {
	Brags []Brag `json:"brags"`
}

type Brag struct {
	Message string `json:"brag"`
}

var brags = AllBragsResponse{}

func init() {
	brags.Brags = []Brag{
		Brag{
			Message: "I have a beautiful fiancé",
		},
	}
}

func main() {
	fmt.Println("new router")
	r := mux.NewRouter()

	fmt.Println("handler")
	r.HandleFunc("/", GetAllBrags).Methods("GET")

	fmt.Println("serving")
	http.ListenAndServe(":8081", r)
}

func GetAllBrags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to get brags")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(brags)
}
