package router

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/darias-developer/test-ms-beer/data"
	"github.com/darias-developer/test-ms-beer/external"
	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/model"
	"github.com/darias-developer/test-ms-beer/service"
	"github.com/gorilla/mux"
)

/* BoxBeerPriceById router: calcula el precio de la caja de cerveza por su id */
func BoxBeerPriceById(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init BoxBeerPriceById")

	params := mux.Vars(r)

	var description string
	var status int
	var id int
	var beerModel model.BeerModel
	var boxpriceRequest data.BoxpriceRequest
	var boxPrice float32
	var liveResponse data.LiveResponse

	err := json.NewDecoder(r.Body).Decode(&boxpriceRequest)

	if err != nil {
		middleware.LogError.Println(err.Error())
		description = "Parametros enviados no son correctos"
		status = http.StatusBadRequest
		goto boxBeerPriceByIdResponse
	}

	if boxpriceRequest.Quantity < 1 {
		boxpriceRequest.Quantity = 6
	}

	id, err = strconv.Atoi(params["id"])

	if err != nil {
		middleware.LogError.Println(err.Error())
		description = "El Id de la cerveza no existe"
		status = http.StatusBadRequest
		goto boxBeerPriceByIdResponse
	}

	beerModel, err = service.FindBeerById(id)

	if err != nil {
		middleware.LogError.Println(err.Error())
		description = "El Id de la cerveza no existe"
		status = http.StatusBadRequest
		goto boxBeerPriceByIdResponse
	}

	middleware.LogInfo.Printf("request id: %d, currency: %s, quantity: %d:", id, boxpriceRequest.Currency, boxpriceRequest.Quantity)

	middleware.LogInfo.Printf("model currency: %s, quantity: %f:", beerModel.Currency, beerModel.Price)

	// en caso de que la la moneda no sea USD se pasa a USD
	if beerModel.Currency != "USD" {

		//valido que el currency sea correcto
		listResponse, err := external.List()

		if err != nil {
			middleware.LogError.Println(err.Error())
			description = "ha ocurrido un error al obtener data desde api.currencylayer.com"
			status = http.StatusInternalServerError
			goto boxBeerPriceByIdResponse
		}

		if listResponse.Currencies[boxpriceRequest.Currency] == "" {
			description = "la moneda ingresada no es valida"
			status = http.StatusNotFound
			goto boxBeerPriceByIdResponse
		}

		/* se agrega un sleep debido a que la version gratuita de api.currencylayer.com
		 * no permite request tan seguidos
		 */
		time.Sleep(2 * time.Second)

		/*
		 * se llama al servicio live para que obtenga el valor en USD de la moneda de la cerveza y
		 * la moneda en que se quiere mostrar el precio de la caja
		 */
		liveResponse, err = external.Live(beerModel.Currency + "," + boxpriceRequest.Currency)

		if err != nil {
			middleware.LogError.Println(err.Error())
			description = "ha ocurrido un error al obtener data desde api.currencylayer.com"
			status = http.StatusInternalServerError
			goto boxBeerPriceByIdResponse
		}

		middleware.LogInfo.Printf("valor 1 USD a %s: %f", beerModel.Currency, liveResponse.Quotes["USD"+beerModel.Currency])

		middleware.LogInfo.Printf("precio original %f en %s", beerModel.Price, beerModel.Currency)

		beerModel.Price = beerModel.Price / float32(liveResponse.Quotes["USD"+beerModel.Currency])

		middleware.LogInfo.Printf("precio transformado %f en USD", beerModel.Price)
	}

	boxPrice = beerModel.Price * float32(boxpriceRequest.Quantity)

	middleware.LogInfo.Printf("precio de caja de cervezas(cantidad %d) en USD: %f", boxpriceRequest.Quantity, boxPrice)

	if boxpriceRequest.Currency != "USD" {
		middleware.LogInfo.Printf("transformar el precio de caja de cervezas a %s", boxpriceRequest.Currency)
		middleware.LogInfo.Printf("valor 1 USD a %s: %f", boxpriceRequest.Currency, liveResponse.Quotes["USD"+boxpriceRequest.Currency])

		boxPrice = boxPrice * float32(liveResponse.Quotes["USD"+boxpriceRequest.Currency])
		middleware.LogInfo.Printf("precio de caja de cervezas(cantidad %d) en %s: %f", boxpriceRequest.Quantity, boxpriceRequest.Currency, boxPrice)
	}

	description = "Operacion exitosa"
	status = http.StatusOK
	goto boxBeerPriceByIdResponse

boxBeerPriceByIdResponse:
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", description)
	rw.WriteHeader(status)
	middleware.LogInfo.Println("end BoxBeerPriceById")
}
