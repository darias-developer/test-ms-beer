package middleware

import (
	"net/http"

	"github.com/darias-developer/test-ms-beer/config"
)

/* CheckDB valida la conexion a la db antes de llamar a una funcion */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		err := config.CheckConn(config.ConnectDB)

		if err != nil {
			rw.Header().Set("Content-Type", "application/json")
			rw.Header().Set("description", "Error en la conexion de la db")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(rw, r)
	}
}
