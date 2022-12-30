package main

import (
	"convenios/configuration"
	"convenios/routes"
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	router := routes.InitRoute()
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	fmt.Println("Inicializando proyecto Convenios")
	if !configuration.Checkconection() {
		log.Fatal("Fallo conexión a BD")
	}
	handleRequests()
}
