package routes

import (
	"github.com/gorilla/mux"
	"github.com/nicolascalvo73/crud-api-rest-go/controllers"
)

func SetUserRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/user/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
}
