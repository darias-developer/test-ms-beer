package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/model"
	"github.com/darias-developer/test-ms-beer/service"
	"github.com/gorilla/mux"
)

/* SearchBeerById router: obtiene una cerveza por medio de su id */
func SearchBeerById(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init SearchBeerById")

	params := mux.Vars(r)
	middleware.LogInfo.Println("id: " + params["id"])

	var description string
	var status int
	var oid string
	var beerModel model.BeerModel

	id, err := strconv.Atoi(params["id"])

	if err != nil {

		middleware.LogError.Println(err.Error())

		description = "El Id de la cerveza no existe"
		status = http.StatusNotFound
	} else {

		beerModel, err = service.FindBeerById(id)

		if err != nil {
			middleware.LogError.Println(err.Error())
			description = "El Id de la cerveza no existe"
			status = http.StatusNotFound
		} else {
			middleware.LogInfo.Println("oid: " + oid)
			description = "Operacion exitosa"
			status = http.StatusOK
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", description)
	rw.WriteHeader(status)

	if status == http.StatusOK {
		json.NewEncoder(rw).Encode(&beerModel)
	}

	middleware.LogInfo.Println("end SearchBeerById")
}
