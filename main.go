package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ianschenck/envflag"

	_ "github.com/lib/pq"
)

// AllBragsResponse is the response with all the brags
type AllBragsResponse struct {
	Brags []Brag `json:"brags"`
}

// Brag represents a brag model
type Brag struct {
	Message string `json:"brag"`
}

type bragDBResult struct {
	bragID  int
	message string
}

var (
	username string
	password string
	host     string
	port     int
	database string
	sslmode  string
)

func init() {
	envflag.StringVar(&username, "USERNAME", "root", "database user")
	envflag.StringVar(&password, "PASSWORD", "sekret", "password for db user")
	envflag.StringVar(&host, "HOST", "localhost", "host where db is running")
	envflag.IntVar(&port, "PORT", 5432, "port the db is running on")
	envflag.StringVar(&database, "DATABASE", "bragsheet", "default database")
	envflag.StringVar(&sslmode, "SSLMODE", "disable", "ssl mode")
}

func main() {
	envflag.Parse()

	log.Printf("host=%s port=%d user=%s dbname=%s sslmode=%s", host, port, username, database, sslmode)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, username, password, database, sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("pinging db: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("pinging db: %v", err)
	}

	fmt.Println("Successfully connected!")
	fmt.Println("new router")
	r := mux.NewRouter()

	log.Println("creating handler funcs")
	r.HandleFunc("/", GetAllBrags(db)).Methods("GET")
	r.HandleFunc("/brag", CreateBrag(db)).Methods("POST")

	fmt.Println("serving")
	http.ListenAndServe(":8080", r)
}

// GetAllBrags fetches all brags
func GetAllBrags(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		fmt.Println("Request to get brags")
		brags := AllBragsResponse{
			Brags: []Brag{},
		}

		rows, err := db.Query("SELECT * FROM brags")
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var bragDB bragDBResult
			err = rows.Scan(&bragDB.bragID, &bragDB.message)

			brag := Brag{
				Message: bragDB.message,
			}

			brags.Brags = append(brags.Brags, brag)
			if err != nil {
				panic(err)
			}
		}
		respondWithJSON(w, 200, brags)
	}
}

// CreateBrag creates a brag
func CreateBrag(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to create brags")

		var brag Brag

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&brag); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		sqlStatement := `INSERT INTO brags (message) VALUES ($1)`
		_, err := db.Exec(sqlStatement, brag.Message)
		if err != nil {
			panic(err)
		}

		brags := AllBragsResponse{
			Brags: []Brag{},
		}
		brags.Brags = append(brags.Brags, brag)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(brags)
	}
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
