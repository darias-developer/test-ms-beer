package external

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/darias-developer/test-ms-beer/data"
)

func Live(currency string) (data.LiveResponse, error) {

	var liveResponse data.LiveResponse

	accessKey := os.Getenv("ACCESS_KEY")

	if accessKey == "" {
		return liveResponse, errors.New("el ACCESS_KEY es requerido")
	}

	if currency == "" {
		return liveResponse, errors.New("el currency es requerido")
	}

	url := "http://api.currencylayer.com/live?access_key=" + accessKey + "&currencies=" + currency + "&source=USD"

	response, err := http.Get(url)

	if err != nil {
		return liveResponse, err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return liveResponse, err
	}

	err = json.Unmarshal(responseData, &liveResponse)

	if err != nil {
		return liveResponse, err
	}

	if !liveResponse.Success {
		return liveResponse, errors.New(liveResponse.Error.Info)
	}

	return liveResponse, nil
}

func List() (data.ListResponse, error) {

	var listResponse data.ListResponse

	accessKey := os.Getenv("ACCESS_KEY")

	if accessKey == "" {
		return listResponse, errors.New("el ACCESS_KEY es requerido")
	}

	url := "http://api.currencylayer.com/list?access_key=" + accessKey

	response, err := http.Get(url)

	if err != nil {
		return listResponse, err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return listResponse, err
	}

	err = json.Unmarshal(responseData, &listResponse)

	if err != nil {
		return listResponse, err
	}

	if !listResponse.Success {
		return listResponse, errors.New(listResponse.Error.Info)
	}

	return listResponse, nil
}
