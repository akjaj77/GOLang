package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type user struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	ID        int64  `json:"id"`
}

type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    user   `json:"employee"`
}

func additem(q http.ResponseWriter, r *http.Request) {

	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)
	q.Header().Set("Content-Type", "application/json")
	profiles = append(profiles, newProfile)
	json.NewEncoder(q).Encode(profiles)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profiles", additem).Methods("POST")
	http.ListenAndServe(":5000", router)

}
