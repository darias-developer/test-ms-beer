package external

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/darias-developer/test-ms-beer/data"
	u "github.com/darias-developer/test-ms-beer/util"
)

type ListType func(utilGet u.TypeGet) (data.ListResponse, error)
type LiveType func(currency string, utilGet u.TypeGet) (data.LiveResponse, error)

func Live(currency string, get u.TypeGet) (data.LiveResponse, error) {

	var liveResponse data.LiveResponse

	accessKey := os.Getenv("ACCESS_KEY")

	if accessKey == "" {
		return liveResponse, errors.New("el ACCESS_KEY es requerido")
	}

	if currency == "" {
		return liveResponse, errors.New("el currency es requerido")
	}

	url := "http://api.currencylayer.com/live?access_key=" + accessKey + "&currencies=" + currency + "&source=USD"

	status, response := get(url)

	if status != http.StatusOK {
		return liveResponse, errors.New(u.ApiCurrencylayerError)
	}

	err := json.Unmarshal(response, &liveResponse)

	if err != nil {
		return liveResponse, err
	}

	if !liveResponse.Success {
		return liveResponse, errors.New(liveResponse.Error.Info)
	}

	return liveResponse, nil
}

func List(get u.TypeGet) (data.ListResponse, error) {

	var listResponse data.ListResponse

	accessKey := os.Getenv("ACCESS_KEY")

	if accessKey == "" {
		return listResponse, errors.New("el ACCESS_KEY es requerido")
	}

	url := "http://api.currencylayer.com/list?access_key=" + accessKey

	status, response := get(url)

	if status != http.StatusOK {
		return listResponse, errors.New(u.ApiCurrencylayerError)
	}

	err := json.Unmarshal(response, &listResponse)

	if err != nil {
		return listResponse, err
	}

	if !listResponse.Success {
		return listResponse, errors.New(listResponse.Error.Info)
	}

	return listResponse, nil
}
