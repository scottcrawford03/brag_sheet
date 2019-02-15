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
	r.HandleFunc("/brag", CreateBrag).Methods("POST")

	fmt.Println("serving")
	http.ListenAndServe(":8080", r)
}

func GetAllBrags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to get brags")

	respondWithJSON(w, 200, brags)
}

func CreateBrag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to create brags")

	var brag Brag

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&brag); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	brags.Brags = append(brags.Brags, brag)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(brags)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
