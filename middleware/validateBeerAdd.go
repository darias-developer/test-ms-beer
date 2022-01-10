package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	m "github.com/darias-developer/test-ms-beer/model"
	u "github.com/darias-developer/test-ms-beer/util"
)

func ValidateBeerAdd(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		original, _ := ioutil.ReadAll(r.Body)
		bufferValidate := ioutil.NopCloser(bytes.NewBuffer(original))
		bufferCopy := ioutil.NopCloser(bytes.NewBuffer(original))

		//valido la data enviada
		status, desc := processBeerAdd(bufferValidate)

		//en caso de error termino la ejecucion del proceso
		if status != http.StatusOK {
			rw.Header().Set("Content-Type", "application/json")
			rw.Header().Set("description", desc)
			rw.WriteHeader(status)
			return
		}

		// en caso de no haber error mantiene el body en el mismo estado y continuo el proceso
		r.Body = bufferCopy

		next.ServeHTTP(rw, r)
	}
}

func processBeerAdd(bufferValidate io.Reader) (int, string) {

	var beerModel m.BeerModel

	err := json.NewDecoder(bufferValidate).Decode(&beerModel)

	//valida que la data enviada sea correcta
	if err != nil {
		u.LogError.Printf("existe un error en la data enviada: %s", err.Error())
		return http.StatusBadRequest, u.BadRequestDesc
	}

	if beerModel.Id == 0 {
		u.LogError.Printf("el campo id es requerido")
		return http.StatusBadRequest, u.BadRequestDesc
	}

	if len(beerModel.Name) == 0 {
		u.LogError.Printf("el campo name es requerido")
		return http.StatusBadRequest, u.BadRequestDesc
	}

	if len(beerModel.Brewery) == 0 {
		u.LogError.Printf("el campo brewery es requerido")
		return http.StatusBadRequest, u.BadRequestDesc
	}

	if len(beerModel.Country) == 0 {
		u.LogError.Printf("el campo country es requerido")
		return http.StatusBadRequest, u.BadRequestDesc
	}

	if beerModel.Price == 0 {
		u.LogError.Printf("el campo price es requerido")
		return http.StatusBadRequest, u.BadRequestDesc
	}

	if len(beerModel.Currency) == 0 {
		u.LogError.Printf("el campo currency es requerido")
		return http.StatusBadRequest, u.BadRequestDesc
	}

	return http.StatusOK, ""
}
