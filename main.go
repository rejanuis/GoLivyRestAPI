package main

import (
	"fmt"
	"log"
	"net/http"

	"./config"
	"./controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	corsorigin := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	router := mux.NewRouter()
	router.HandleFunc("/runlivy", controller.RunLivy).Methods("POST")

	//add swagger folder to display list service
	swagger := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/swaggerui/").Handler(swagger)
	http.Handle("/", corsorigin.Handler(router))

	fmt.Println("Connected to port " + config.GetConfig("portapi"))
	log.Fatal(http.ListenAndServe(":"+config.GetConfig("portapi"), nil))

}
