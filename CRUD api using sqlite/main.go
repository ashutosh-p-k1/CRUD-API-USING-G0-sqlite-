package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Policy struct {
	APIVersion string `json:"apiversion"`
	Kind       string `json:"kind"`
	Metadata   string `json:"metadata"`
	Spec       string `json:"spec"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/create", createhandler).Methods("POST")
	r.HandleFunc("/read", readhandler).Methods("GET")
	r.HandleFunc("/delete/{id}", deleteHandler).Methods("DELETE")
	r.HandleFunc("/update/{id}", updateHandler).Methods("POST")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}

func createhandler(w http.ResponseWriter, r *http.Request) {
	createPolicy := Policy{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in Reading the body :=>", err)
	}
	err = json.Unmarshal([]byte(body), &createPolicy)
	if err != nil {
		log.Println("Error in unmarshalling :=>", err)
	}
	db := createDatabase()
	insertDatabase(createPolicy, db)
}

func readhandler(w http.ResponseWriter, r *http.Request) {
	db := createDatabase()
	newpolicies := Readdatabase(db)

	res, _ := json.Marshal(newpolicies)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	db := createDatabase()
	DeleteDB(ids, db)

}
func updateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Inside Update")
	updatePolicy := Policy{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in Reading the body :=>", err)
	}
	err = json.Unmarshal([]byte(body), &updatePolicy)

	if err != nil {
		log.Println("Error in unmarshalling :=>", err)
	}
	params := mux.Vars(r)
	ids := params["id"]
	db := createDatabase()
	UpdateDB(updatePolicy, ids, db)
}
