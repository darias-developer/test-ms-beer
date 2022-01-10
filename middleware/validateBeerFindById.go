package middleware

import (
	"net/http"
	"strconv"

	u "github.com/darias-developer/test-ms-beer/util"
	"github.com/gorilla/mux"
)

func ValidateBeerFindById(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		//valido la data enviada
		status, desc := processBeerFindById(r)

		//en caso de error termino la ejecucion del proceso
		if status != http.StatusOK {
			rw.Header().Set("Content-Type", "application/json")
			rw.Header().Set("description", desc)
			rw.WriteHeader(status)
			return
		}

		next.ServeHTTP(rw, r)
	}
}

func processBeerFindById(r *http.Request) (int, string) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.LogError.Println(err.Error())
		return http.StatusBadRequest, u.BeerIdError
	}

	u.LogInfo.Printf("id: %v", id)

	return http.StatusOK, ""
}
