package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/darias-developer/test-ms-beer/controller"
	"github.com/darias-developer/test-ms-beer/data"
	"github.com/darias-developer/test-ms-beer/external"
	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/model"
	"github.com/darias-developer/test-ms-beer/service"
	"github.com/gorilla/mux"
)

/* BeerAdd router: agrega cervezas */
func BeerAdd(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init BeerAdd")

	var status int = http.StatusBadRequest
	var desc string = "Request invalida"

	var beerModel model.BeerModel
	err := json.NewDecoder(r.Body).Decode(&beerModel)

	if err != nil {
		middleware.LogError.Printf("existe un error en la data enviada: %s", err.Error())
	} else {
		status, desc = controller.BeerAdd(service.BeerAddService, service.BeerFindByIdService, external.List, beerModel)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	middleware.LogInfo.Println("end BeerAdd")
}

/* BeerFindAll router: lista todas las cervezas */
func BeerFindAll(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init BeerFindAll")

	status, desc, arr := controller.BeerFindAll(service.BeerFindAllService)

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	if len(arr) > 0 {
		json.NewEncoder(rw).Encode(&arr)
	}

	middleware.LogInfo.Println("end BeerFindAll")
}

/* BeerFindById router: obtiene una cerveza por medio de su id */
func BeerFindById(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init BeerFindById")

	var status int = http.StatusBadRequest
	var desc string = "Request invalida"
	var beerModel model.BeerModel

	params := mux.Vars(r)
	middleware.LogInfo.Println("id: " + params["id"])

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		middleware.LogError.Println(err.Error())
		desc = "El Id de la cerveza no es valido"
		status = http.StatusNotFound
	} else {
		status, desc, beerModel = controller.BeerFindById(service.BeerFindByIdService, id)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	if status == http.StatusOK {
		json.NewEncoder(rw).Encode(&beerModel)
	}

	middleware.LogInfo.Println("end BeerFindById")
}

/* BeerBoxPriceById router: calcula el precio de la caja de cerveza por su id */
func BeerBoxPriceById(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init BeerBoxPriceById")

	var status int
	var desc string
	var boxPrice float32

	params := mux.Vars(r)

	middleware.LogInfo.Println("id: " + params["id"])

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		middleware.LogError.Println(err.Error())
		desc = "El Id de la cerveza no es valido"
		status = http.StatusNotFound
	} else {
		var boxpriceRequest data.BoxpriceRequest
		err := json.NewDecoder(r.Body).Decode(&boxpriceRequest)

		if err != nil {
			middleware.LogError.Printf("existe un error en la data enviada: %s", err.Error())
			status = http.StatusBadRequest
			desc = "Request invalida"
		} else {
			status, desc, boxPrice = controller.BeerBoxPriceById(
				service.BeerFindByIdService, external.List, external.Live, boxpriceRequest, id)
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	if status == http.StatusOK {
		json.NewEncoder(rw).Encode(&data.BoxpriceResponse{PriceTotal: boxPrice})
	}

	middleware.LogInfo.Println("end BeerBoxPriceById")
}
