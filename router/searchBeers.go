package router

import (
	"encoding/json"
	"net/http"

	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/service"
)

/* SearchBeers router: lista todas las cervezas */
func SearchBeers(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init SearchBeers")

	arr, err := service.SearchBeers()

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", "Operacion exitosa")
	rw.WriteHeader(http.StatusOK)

	if err != nil {
		middleware.LogError.Println(err.Error())
	} else {
		json.NewEncoder(rw).Encode(&arr)
	}

	middleware.LogInfo.Println("end SearchBeers")
}
