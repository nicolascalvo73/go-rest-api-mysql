package commons

import "net/http"

func SendResponse(writer http.ResponseWriter, status int, data []byte) { //w=write respuesta
	writer.Header().Set("Content-Type", "application/json") //escribe lo que devuelve la funcion a sabiendas que el contanido es JSON
	writer.WriteHeader(status)
	writer.Write(data)
}

func SendError(writer http.ResponseWriter, status int) {
	data := []byte(`{}`)                                    //variable de slice de bytes
	writer.Header().Set("Content-Type", "application/json") //asignamos tipo JSON para que no devuelva int
	writer.WriteHeader(status)                              //en caso de error enviara un JSON vacio
	writer.Write(data)
}
