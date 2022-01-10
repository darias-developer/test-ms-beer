package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	d "github.com/darias-developer/test-ms-beer/data"
	u "github.com/darias-developer/test-ms-beer/util"
	"github.com/gorilla/mux"
)

func ValidateBeerBoxPriceById(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		original, _ := ioutil.ReadAll(r.Body)
		bufferValidate := ioutil.NopCloser(bytes.NewBuffer(original))
		bufferCopy := ioutil.NopCloser(bytes.NewBuffer(original))

		//valido la data enviada
		status, desc := processBeerBoxPriceById(bufferValidate, r)

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

func processBeerBoxPriceById(bufferValidate io.Reader, r *http.Request) (int, string) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.LogError.Println(err.Error())
		return http.StatusBadRequest, u.BeerIdError
	}

	u.LogInfo.Printf("id: %v", id)

	var boxpriceRequest d.BoxpriceRequest

	err = json.NewDecoder(bufferValidate).Decode(&boxpriceRequest)

	//valida que la data enviada sea correcta
	if err != nil {
		u.LogError.Printf("existe un error en la data enviada: %s", err.Error())
		return http.StatusBadRequest, u.BadRequestDesc
	}

	if len(boxpriceRequest.Currency) == 0 {
		u.LogError.Printf("el campo currency es requerido")
		return http.StatusBadRequest, u.BadRequestDesc
	}

	return http.StatusOK, ""
}
