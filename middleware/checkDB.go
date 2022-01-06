package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/darias-developer/test-ms-beer/config"
	"github.com/darias-developer/test-ms-beer/data"
)

var response data.Response

/* CheckDB valida la conexion a la db antes de llamar a una funcion */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		response = data.Response{
			ResponseCode: "ERROR",
			Description:  "Error en la conexion de la db",
		}

		responseJson, _ := json.Marshal(response)

		if config.CheckConnection() == 0 {
			http.Error(rw, string(responseJson), 500)
		}

		next.ServeHTTP(rw, r)
	}
}
