package controller

import (
	"net/http"
	"time"

	"github.com/darias-developer/test-ms-beer/data"
	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/model"
)

type BeerAddService func(beerModel model.BeerModel) (string, error)
type BeerFindByIdService func(id int) (model.BeerModel, error)
type BeerFindAllService func() ([]model.BeerModel, error)
type ListExternal func() (data.ListResponse, error)
type LiveExternal func(currencies string) (data.LiveResponse, error)

var descBadRequest string = "Request invalida"
var descInternalServerError string = "Ha ocurrido un error"
var descConflict string = "El ID de la cerveza ya existe"

func BeerAdd(beerAddService BeerAddService, beerFindByIdService BeerFindByIdService,
	listExternal ListExternal, beerModel model.BeerModel) (int, string) {

	if beerModel.Id == 0 {
		middleware.LogError.Printf("el campo id es requerido")
		return http.StatusBadRequest, descBadRequest
	}

	_, err := beerFindByIdService(beerModel.Id)

	if err == nil {
		return http.StatusConflict, descConflict
	}

	middleware.LogInfo.Println("Id no registrado. Se procede a registrar la cerveza")

	if len(beerModel.Name) == 0 {
		middleware.LogError.Printf("el campo name es requerido")
		return http.StatusBadRequest, descBadRequest
	}

	if len(beerModel.Brewery) == 0 {
		middleware.LogError.Printf("el campo brewery es requerido")
		return http.StatusBadRequest, descBadRequest
	}

	if len(beerModel.Country) == 0 {
		middleware.LogError.Printf("el campo country es requerido")
		return http.StatusBadRequest, descBadRequest
	}

	if beerModel.Price == 0 {
		middleware.LogError.Printf("el campo price es requerido")
		return http.StatusBadRequest, descBadRequest
	}

	if len(beerModel.Currency) == 0 {
		middleware.LogError.Printf("el campo currency es requerido")
		return http.StatusBadRequest, descBadRequest
	}

	//valido que el currency sea correcto
	listResponse, err := listExternal()

	if err != nil {
		middleware.LogError.Printf("ha ocurrido un error al obtener data desde api.currencylayer.com: %s", err.Error())
		return http.StatusBadRequest, descBadRequest
	}

	if listResponse.Currencies[beerModel.Currency] == "" {
		return http.StatusBadRequest, descBadRequest
	}

	oid, err := beerAddService(beerModel)

	if err != nil {
		middleware.LogError.Println(err)
		return http.StatusInternalServerError, "Ha ocurrido un error al crear la cerveza"
	}

	middleware.LogInfo.Printf("oid creado: %s", oid)

	return http.StatusOK, "Cerveza creada"
}

func BeerFindAll(beerFindAllService BeerFindAllService) (int, string, []model.BeerModel) {

	arr, err := beerFindAllService()

	if err != nil {
		middleware.LogError.Println(err.Error())
	}

	middleware.LogInfo.Printf("registros encontrados: %v", len(arr))

	return http.StatusOK, "Cerveza creada", arr
}

func BeerFindById(beerFindByIdService BeerFindByIdService, id int) (int, string, model.BeerModel) {

	beerModel, err := beerFindByIdService(id)

	if err != nil {
		return http.StatusNotFound, "El Id de la cerveza no existe", beerModel
	}

	middleware.LogInfo.Printf("oid: %s", beerModel.OID.Hex())

	return http.StatusOK, "Operacion exitosa", beerModel
}

func BeerBoxPriceById(beerFindByIdService BeerFindByIdService, listExternal ListExternal,
	liveExternal LiveExternal, boxpriceRequest data.BoxpriceRequest, id int) (int, string, float32) {

	// en caso de que el campo Quantity no sea valido pór defecto queda en 6
	if boxpriceRequest.Quantity < 1 {
		boxpriceRequest.Quantity = 6
	}

	beerModel, err := beerFindByIdService(id)

	if err != nil {
		return http.StatusNotFound, "El Id de la cerveza no existe", 0
	}

	middleware.LogInfo.Printf("request id: %d, currency: %s, quantity: %d:", id, boxpriceRequest.Currency, boxpriceRequest.Quantity)

	middleware.LogInfo.Printf("model currency: %s, quantity: %f:", beerModel.Currency, beerModel.Price)

	var liveResponse data.LiveResponse

	//valido que el currency sea correcto
	listResponse, err := listExternal()

	if err != nil {
		middleware.LogError.Printf("%v: %s", err.Error(),
			"ha ocurrido un error al obtener data desde api.currencylayer.com")
		return http.StatusInternalServerError, descInternalServerError, 0
	}

	if listResponse.Currencies[boxpriceRequest.Currency] == "" {
		middleware.LogError.Printf("la moneda ingresada no es valida")
		return http.StatusNotFound, descBadRequest, 0
	}

	// en caso de que la la moneda no sea USD se pasa a USD
	if beerModel.Currency != "USD" {

		/* se agrega un sleep debido a que la version gratuita de api.currencylayer.com
		 * no permite request tan seguidos
		 */
		time.Sleep(2 * time.Second)

		/*
		 * se llama al servicio live para que obtenga el valor en USD de la moneda de la cerveza y
		 * la moneda en que se quiere mostrar el precio de la caja
		 */
		liveResponse, err = liveExternal(beerModel.Currency + "," + boxpriceRequest.Currency)

		if err != nil {
			middleware.LogError.Printf("%v: %s", err.Error(),
				"ha ocurrido un error al obtener data desde api.currencylayer.com")
			return http.StatusInternalServerError, descInternalServerError, 0
		}

		middleware.LogInfo.Printf("valor 1 USD a %s: %f", beerModel.Currency, liveResponse.Quotes["USD"+beerModel.Currency])

		middleware.LogInfo.Printf("precio original %f en %s", beerModel.Price, beerModel.Currency)

		beerModel.Price = beerModel.Price / float32(liveResponse.Quotes["USD"+beerModel.Currency])

		middleware.LogInfo.Printf("precio transformado %f en USD", beerModel.Price)
	}

	boxPrice := beerModel.Price * float32(boxpriceRequest.Quantity)

	middleware.LogInfo.Printf("precio de caja de cervezas(cantidad %d) en USD: %f", boxpriceRequest.Quantity, boxPrice)

	if boxpriceRequest.Currency != "USD" {
		middleware.LogInfo.Printf("transformar el precio de caja de cervezas a %s", boxpriceRequest.Currency)
		middleware.LogInfo.Printf("valor 1 USD a %s: %f", boxpriceRequest.Currency, liveResponse.Quotes["USD"+boxpriceRequest.Currency])

		boxPrice = boxPrice * float32(liveResponse.Quotes["USD"+boxpriceRequest.Currency])
		middleware.LogInfo.Printf("precio de caja de cervezas(cantidad %d) en %s: %f", boxpriceRequest.Quantity, boxpriceRequest.Currency, boxPrice)
	}

	return http.StatusOK, "Operacion exitosa", boxPrice
}
