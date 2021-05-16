package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolascalvo73/crud-api-rest-go/commons"
	"github.com/nicolascalvo73/crud-api-rest-go/routes"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()

	routes.SetUserRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")
	log.Println(server.ListenAndServe())

}
