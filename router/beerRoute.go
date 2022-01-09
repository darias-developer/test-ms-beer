package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	c "github.com/darias-developer/test-ms-beer/controller"
	d "github.com/darias-developer/test-ms-beer/data"
	e "github.com/darias-developer/test-ms-beer/external"
	m "github.com/darias-developer/test-ms-beer/model"
	s "github.com/darias-developer/test-ms-beer/service"
	u "github.com/darias-developer/test-ms-beer/util"
	"github.com/gorilla/mux"
)

/* BeerAdd router: agrega cervezas */
func BeerAdd(rw http.ResponseWriter, r *http.Request) {

	u.LogInfo.Println("init BeerAdd")

	var status int = http.StatusBadRequest
	var desc string = "Request invalida"

	var beerModel m.BeerModel
	err := json.NewDecoder(r.Body).Decode(&beerModel)

	if err != nil {
		u.LogError.Printf("existe un error en la data enviada: %s", err.Error())
	} else {
		status, desc = c.BeerAdd(s.BeerAdd, s.BeerFindById, e.List, beerModel)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	u.LogInfo.Println("end BeerAdd")
}

/* BeerFindAll router: lista todas las cervezas */
func BeerFindAll(rw http.ResponseWriter, r *http.Request) {

	u.LogInfo.Println("init BeerFindAll")

	status, desc, arr := c.BeerFindAll(s.BeerFindAll)

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	if len(arr) > 0 {
		json.NewEncoder(rw).Encode(&arr)
	}

	u.LogInfo.Println("end BeerFindAll")
}

/* BeerFindById router: obtiene una cerveza por medio de su id */
func BeerFindById(rw http.ResponseWriter, r *http.Request) {

	u.LogInfo.Println("init BeerFindById")

	var status int = http.StatusBadRequest
	var desc string = "Request invalida"
	var beerModel m.BeerModel

	params := mux.Vars(r)
	u.LogInfo.Println("id: " + params["id"])

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.LogError.Println(err.Error())
		desc = "El Id de la cerveza no es valido"
		status = http.StatusNotFound
	} else {
		status, desc, beerModel = c.BeerFindById(s.BeerFindById, id)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	if status == http.StatusOK {
		json.NewEncoder(rw).Encode(&beerModel)
	}

	u.LogInfo.Println("end BeerFindById")
}

/* BeerBoxPriceById router: calcula el precio de la caja de cerveza por su id */
func BeerBoxPriceById(rw http.ResponseWriter, r *http.Request) {

	u.LogInfo.Println("init BeerBoxPriceById")

	var status int
	var desc string
	var boxPrice float32

	params := mux.Vars(r)

	u.LogInfo.Println("id: " + params["id"])

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.LogError.Println(err.Error())
		desc = "El Id de la cerveza no es valido"
		status = http.StatusNotFound
	} else {
		var boxpriceRequest d.BoxpriceRequest
		err := json.NewDecoder(r.Body).Decode(&boxpriceRequest)

		if err != nil {
			u.LogError.Printf("existe un error en la data enviada: %s", err.Error())
			status = http.StatusBadRequest
			desc = "Request invalida"
		} else {
			status, desc, boxPrice = c.BeerBoxPriceById(s.BeerFindById, e.List, e.Live, boxpriceRequest, id)
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	if status == http.StatusOK {
		json.NewEncoder(rw).Encode(&d.BoxpriceResponse{PriceTotal: boxPrice})
	}

	u.LogInfo.Println("end BeerBoxPriceById")
}
