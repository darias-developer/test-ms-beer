package controller

import (
	"net/http"

	"github.com/darias-developer/test-ms-beer/data"
	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/model"
)

type CreateBeerService func(beerModel model.BeerModel) (string, error)
type FindBeerByIdService func(id int) (model.BeerModel, error)
type BeerFindAllService func() ([]*model.BeerModel, error)
type ListExternal func() (data.ListResponse, error)

var descBadRequest string = "Request invalida"
var descConflict string = "El ID de la cerveza ya existe"

func BeerAdd(createBeerService CreateBeerService, findBeerByIdService FindBeerByIdService,
	listExternal ListExternal, beerModel model.BeerModel) (int, string) {

	if beerModel.Id == 0 {
		middleware.LogError.Printf("el campo id es requerido")
		return http.StatusBadRequest, descBadRequest
	}

	_, err := findBeerByIdService(beerModel.Id)

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

	oid, err := createBeerService(beerModel)

	if err != nil {
		middleware.LogError.Println(err)
		return http.StatusInternalServerError, "Ha ocurrido un error al crear la cerveza"
	}

	middleware.LogInfo.Printf("oid creado: %s", oid)

	return http.StatusOK, "Cerveza creada"
}

func BeerFindAll(beerFindAllService BeerFindAllService) (int, string, []*model.BeerModel) {

	arr, err := beerFindAllService()

	if err != nil {
		middleware.LogError.Println(err.Error())
	}

	return http.StatusOK, "Cerveza creada", arr
}
