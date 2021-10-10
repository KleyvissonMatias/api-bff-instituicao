package main

import (
	router "api-bff-instituicao/controller"
	service "api-bff-instituicao/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var conn = service.ServiceDb{}

func init() {
	conn.Connect()
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/instituicoes", router.Create).Methods("POST")
	r.HandleFunc("/api/v1/instituicoes/{id}", router.Update).Methods("PUT")
	r.HandleFunc("/api/v1/instituicoes", router.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/instituicoes/{id}", router.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/instituicoes{id}", router.Delete).Methods("DELETE")

	port := ":3000"
	fmt.Println("Running in port: ", port)
	log.Fatal(http.ListenAndServe(port, r))
}
