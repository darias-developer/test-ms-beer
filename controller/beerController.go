package controller

import (
	"net/http"
	"time"

	"github.com/darias-developer/test-ms-beer/config"
	d "github.com/darias-developer/test-ms-beer/data"
	e "github.com/darias-developer/test-ms-beer/external"
	m "github.com/darias-developer/test-ms-beer/model"
	s "github.com/darias-developer/test-ms-beer/service"
	u "github.com/darias-developer/test-ms-beer/util"
)

func BeerAdd(add s.BeerAddType, findBy s.BeerFindByIdType, list e.ListType, beer m.BeerModel) (int, string) {

	//valido que el id ingresado no exista en la db
	_, err := findBy(beer.Id, config.ConnectDB, u.FindOneBeer)

	if err == nil {
		return http.StatusConflict, u.BeerNotFound
	}

	u.LogInfo.Println("Id no registrado. Se procede a registrar la cerveza")

	//valido que el currency sea correcto
	listResponse, err := list(u.Get)

	if err != nil {
		u.LogError.Printf(err.Error())
		return http.StatusBadRequest, u.BadRequestDesc
	}

	if listResponse.Currencies[beer.Currency] == "" {
		return http.StatusBadRequest, u.BadRequestDesc
	}

	oid, err := add(beer, config.ConnectDB, u.InsertOneBeer)

	if err != nil {
		u.LogError.Println(err)
		return http.StatusInternalServerError, u.BeerError
	}

	u.LogInfo.Printf("oid creado: %s", oid)

	return http.StatusOK, u.BeerCreated
}

func BeerFindAll(beerFindAll s.BeerFindAllType) (int, string, []m.BeerModel) {

	arr, err := beerFindAll(config.ConnectDB, u.FindAllBeer)

	if err != nil {
		u.LogError.Println(err.Error())
	}

	u.LogInfo.Printf("registros encontrados: %v", len(arr))

	return http.StatusOK, u.SuccessOperation, arr
}

func BeerFindById(findBy s.BeerFindByIdType, id int) (int, string, m.BeerModel) {

	beerModel, err := findBy(id, config.ConnectDB, u.FindOneBeer)

	if err != nil {
		return http.StatusNotFound, u.BeerNotFound, beerModel
	}

	u.LogInfo.Printf("oid: %s", beerModel.OID.Hex())

	return http.StatusOK, u.SuccessOperation, beerModel
}

func BeerBoxPriceById(
	findBy s.BeerFindByIdType, list e.ListType, live e.LiveType, boxpriceReq d.BoxpriceRequest, id int) (int, string, float32) {

	// en caso de que el campo Quantity no sea valido p??r defecto queda en 6
	if boxpriceReq.Quantity < 1 {
		boxpriceReq.Quantity = 6
	}

	beer, err := findBy(id, config.ConnectDB, u.FindOneBeer)

	if err != nil {
		return http.StatusNotFound, u.BeerNotFound, 0
	}

	u.LogInfo.Printf("request id: %d, currency: %s, quantity: %d:", id, boxpriceReq.Currency, boxpriceReq.Quantity)

	u.LogInfo.Printf("model currency: %s, quantity: %f:", beer.Currency, beer.Price)

	//valido que el currency sea correcto
	listResponse, err := list(u.Get)

	if err != nil {
		u.LogError.Println(err.Error())
		return http.StatusInternalServerError, u.InternalServerErrorDesc, 0
	}

	if listResponse.Currencies[boxpriceReq.Currency] == "" {
		u.LogError.Printf("la moneda ingresada no es valida")
		return http.StatusNotFound, u.BadRequestDesc, 0
	}

	var liveResponse d.LiveResponse

	// en caso de que la la moneda no sea USD se pasa a USD
	if beer.Currency != "USD" {

		/* se agrega un sleep debido a que la version gratuita de api.currencylayer.com
		 * no permite request tan seguidos
		 */
		time.Sleep(2 * time.Second)

		/*
		 * se llama al servicio live para que obtenga el valor en USD de la moneda de la cerveza y
		 * la moneda en que se quiere mostrar el precio de la caja
		 */

		currencies := beer.Currency + "," + boxpriceReq.Currency

		liveResponse, err = live(currencies, u.Get)

		if err != nil {
			u.LogError.Println(err.Error())
			return http.StatusInternalServerError, u.InternalServerErrorDesc, 0
		}

		u.LogInfo.Printf("valor 1 USD a %s: %f", beer.Currency, liveResponse.Quotes["USD"+beer.Currency])

		u.LogInfo.Printf("precio original %f en %s", beer.Price, beer.Currency)

		beer.Price = beer.Price / float32(liveResponse.Quotes["USD"+beer.Currency])

		u.LogInfo.Printf("precio transformado %f en USD", beer.Price)
	}

	boxPrice := beer.Price * float32(boxpriceReq.Quantity)

	u.LogInfo.Printf("precio de caja de cervezas(cantidad %d) en USD: %f", boxpriceReq.Quantity, boxPrice)

	if boxpriceReq.Currency != "USD" {
		u.LogInfo.Printf("transformar el precio de caja de cervezas a %s", boxpriceReq.Currency)
		u.LogInfo.Printf("valor 1 USD a %s: %f", boxpriceReq.Currency, liveResponse.Quotes["USD"+boxpriceReq.Currency])

		boxPrice = boxPrice * float32(liveResponse.Quotes["USD"+boxpriceReq.Currency])
		u.LogInfo.Printf("precio de caja de cervezas(cantidad %d) en %s: %f", boxpriceReq.Quantity, boxpriceReq.Currency, boxPrice)
	}

	return http.StatusOK, u.SuccessOperation, boxPrice
}
