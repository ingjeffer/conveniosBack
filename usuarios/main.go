package main

import (
	"fmt"
	"log"
	"net/http"
	"usuarios/configuration"
	"usuarios/routes"
)

func handleRequests() {
	router := routes.InitRoute()
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Inicializando proyecto Usuarios")
	configuration.IniciarDB()
	configuration.IniciarRedis()
	handleRequests()
}
