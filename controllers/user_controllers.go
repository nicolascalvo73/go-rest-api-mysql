package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolascalvo73/crud-api-rest-go/commons"
	"github.com/nicolascalvo73/crud-api-rest-go/models"
)

func GetAll(writer http.ResponseWriter, request *http.Request) { //el segundo valor es un pointer
	users := []models.User{}
	db := commons.GetConnections()
	defer db.Close()

	db.Find(&users)

	json, _ := json.Marshal(users)                    //el "_" sirve para ignorar el error que devuelve la función marshall
	commons.SendResponse(writer, http.StatusOK, json) //(writer, http.StatusOK, json) son los argunmentos que le pasa a la func Send Response
}

func Get(writer http.ResponseWriter, request *http.Request) { //func para acceder a un solo user
	user := models.User{}

	id := mux.Vars(request)["ID"]
	db := commons.GetConnections()
	defer db.Close()

	db.Find(&user, id)

	if user.ID > 0 {
		json, _ := json.Marshal(user)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	db := commons.GetConnections()
	defer db.Close()

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	err = db.Save(&user).Error
	if err != nil {
		log.Fatal(err)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(user)
	commons.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	//id := mux.Vars(request)["ID"]
	db := commons.GetConnections()
	defer db.Close()
	db.Find(&user.ID)

	if user.ID > 0 {
		db.Delete(user)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
