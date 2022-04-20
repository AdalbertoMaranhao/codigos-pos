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
		Title:       "apis",
		Description: "Testando os eventos",
	},
}

// @title           Swagger API-Event
// @version         1.0
// @description     Documentação da API de eventos.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API-Event Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://new-api-events-16041.herokuapp.com/

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

// ShowAllEvents godoc
// @Summary      Show all events
// @Description  List all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Success      200  {object}  Event
// @Router       /events [get]

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("acessando o endpoint get all events")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
