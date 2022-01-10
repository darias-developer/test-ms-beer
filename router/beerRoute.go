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

	//transforma el json a un objecto
	var beerModel m.BeerModel
	json.NewDecoder(r.Body).Decode(&beerModel)

	status, desc := c.BeerAdd(s.BeerAdd, s.BeerFindById, e.List, beerModel)

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

	//obtierne el parametro id desde la url
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	status, desc, beerModel := c.BeerFindById(s.BeerFindById, id)

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

	//transforma el json a un objecto
	var boxpriceRequest d.BoxpriceRequest
	json.NewDecoder(r.Body).Decode(&boxpriceRequest)

	//obtierne el parametro id desde la url
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	status, desc, boxPrice := c.BeerBoxPriceById(s.BeerFindById, e.List, e.Live, boxpriceRequest, id)

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", desc)
	rw.WriteHeader(status)

	if status == http.StatusOK {
		json.NewEncoder(rw).Encode(&d.BoxpriceResponse{PriceTotal: boxPrice})
	}

	u.LogInfo.Println("end BeerBoxPriceById")
}
