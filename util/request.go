package util

import (
	"io/ioutil"
	"net/http"
)

type TypeGet func(url string) (int, []byte)

func Get(url string) (int, []byte) {

	LogInfo.Printf("url: %s", url)

	response, err := http.Get(url)
	LogInfo.Printf("response status : %v", response.StatusCode)

	if err != nil {
		LogError.Printf(err.Error())
		return response.StatusCode, nil
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		LogError.Printf(err.Error())
		return response.StatusCode, nil
	}

	return response.StatusCode, responseData
}
