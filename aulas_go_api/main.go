package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Event struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []Event

var events = allEvents{
	{
		Id:          1,
		Title:       "APIs",
		Description: "Testando os eventos",
	},
}

func main() {
	log.Println("Starting API")
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/health-check", HealtCheck).Methods("GET")
	router.HandleFunc("/events", GetAllEvents).Methods("GET")

	http.ListenAndServe(":"+port, router)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func HealtCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando health-check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aplicação em Funcionando!\n")
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("acessando o endpoint get all events")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
