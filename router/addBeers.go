package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/darias-developer/test-ms-beer/external"
	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/model"
	"github.com/darias-developer/test-ms-beer/service"
)

/* AddBeers router: agrega cervezas */
func AddBeers(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init AddBeers")

	status, err := process(r)

	var description string

	if err != nil {

		middleware.LogError.Println(err.Error())

		if status == http.StatusBadRequest {
			description = "Request invalida"
		} else {
			description = err.Error()
		}

	} else {

		description = "Cerveza creada"
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("description", description)
	rw.WriteHeader(status)

	middleware.LogInfo.Println("end AddBeers")
}

func process(r *http.Request) (int, error) {

	var beerModel model.BeerModel
	err := json.NewDecoder(r.Body).Decode(&beerModel)

	if err != nil {
		middleware.LogError.Println(err)
		return http.StatusBadRequest, errors.New("existe un error en la data enviada")
	}

	if beerModel.Id == 0 {
		return http.StatusBadRequest, errors.New("el campo id es requerido")
	}

	_, err = service.FindBeerById(beerModel.Id)

	if err != nil {
		middleware.LogWarn.Println(err)
		middleware.LogInfo.Println("Id no registrado. Se procede a registrar la cerveza")
	} else {
		return http.StatusConflict, errors.New("El ID de la cerveza ya existe")
	}

	if len(beerModel.Name) == 0 {
		return http.StatusBadRequest, errors.New("el campo name es requerido")
	}

	if len(beerModel.Brewery) == 0 {
		return http.StatusBadRequest, errors.New("el campo brewery es requerido")
	}

	if len(beerModel.Country) == 0 {
		return http.StatusBadRequest, errors.New("el campo country es requerido")
	}

	if beerModel.Price == 0 {
		return http.StatusBadRequest, errors.New("el campo price es requerido")
	}

	if len(beerModel.Currency) == 0 {
		return http.StatusBadRequest, errors.New("el campo currency es requerido")
	}

	//valido que el currency sea correcto
	listResponse, err := external.List()

	if err != nil {
		middleware.LogError.Println(err.Error())
		return http.StatusBadRequest, errors.New("ha ocurrido un error al obtener data desde api.currencylayer.com")
	}

	if listResponse.Currencies[beerModel.Currency] == "" {
		return http.StatusBadRequest, errors.New("la moneda ingresada no es valida")
	}

	oid, err := service.CreateBeer(beerModel)

	if err != nil {
		middleware.LogError.Println(err)
		return http.StatusInternalServerError, errors.New("ha ocurrido un error al crear la cerveza")
	}

	middleware.LogInfo.Println("oid creado: " + oid)

	return http.StatusCreated, nil
}
