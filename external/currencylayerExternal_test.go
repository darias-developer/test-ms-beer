package external

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/darias-developer/test-ms-beer/data"
)

func liveResponse() data.LiveResponse {

	var liveResponse data.LiveResponse

	m := make(map[string]float32)

	m["USDEUR"] = 0.884795
	m["USDUSD"] = 1
	m["USDCLP"] = 837.598714

	liveResponse.Success = true
	liveResponse.Quotes = m

	return liveResponse
}

func listResponse() data.ListResponse {

	var listResponse data.ListResponse

	m := make(map[string]string)
	m["USD"] = "USD"
	m["EUR"] = "EUR"
	m["CLP"] = "CLP"

	listResponse.Success = true
	listResponse.Currencies = m

	return listResponse
}

func mock_live_success(url string) (int, []byte) {

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(liveResponse())

	return 200, reqBodyBytes.Bytes()
}

func mock_live_success_with_error(url string) (int, []byte) {
	return 200, nil
}

func mock_live_success_with_data_error(url string) (int, []byte) {

	liveResponse := liveResponse()
	liveResponse.Success = false
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(liveResponse)

	return 200, reqBodyBytes.Bytes()
}

func mock_live_error(url string) (int, []byte) {
	return 500, nil
}

func mock_list_success(url string) (int, []byte) {

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(listResponse())

	return 200, reqBodyBytes.Bytes()
}

func mock_list_success_with_error(url string) (int, []byte) {
	return 200, nil
}

func mock_list_success_with_data_error(url string) (int, []byte) {

	listResponse := listResponse()
	listResponse.Success = false
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(listResponse)

	return 200, reqBodyBytes.Bytes()
}

func mock_list_error(url string) (int, []byte) {
	return 500, nil
}

func TestLive(t *testing.T) {

	t.Run("test Live valida el parametro currency", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		_, err := Live("", mock_live_success)

		if err.Error() != "el currency es requerido" {
			t.Errorf("Expected: %s, got: %s", "el currency es requerido", err.Error())
		}
	})

	t.Run("test Live valida el parametro ACCESS_KEY", func(t *testing.T) {

		_, err := Live("USD", mock_live_success)

		if err.Error() != "el ACCESS_KEY es requerido" {
			t.Errorf("Expected: %v, got: %v", "el ACCESS_KEY es requerido", err.Error())
		}
	})

	t.Run("test Live error al accesar los servicios externos", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		_, err := Live("USD", mock_live_error)

		if err.Error() != "ha ocurrio un error en los servicios api.currencylayer.com" {
			t.Errorf("Expected: %v, got: %v", "ha ocurrio un error en los servicios api.currencylayer.com", err.Error())
		}
	})

	t.Run("test Live error al parsear la respuesta del servicio", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		_, err := Live("USD", mock_live_success_with_error)

		if err.Error() != "unexpected end of JSON input" {
			t.Errorf("Expected: %v, got: %v", "unexpected end of JSON input", err.Error())
		}
	})

	t.Run("test Live error en la respuesta", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		response, _ := Live("USD", mock_live_success_with_data_error)

		if response.Success {
			t.Errorf("Expected: %v, got: %v", false, response.Success)
		}
	})

	t.Run("test Live success", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		_, err := Live("USD", mock_live_success)

		if err != nil {
			t.Errorf("Expected: %v, got: %v", nil, err.Error())
		}
	})
}

func TestList(t *testing.T) {

	t.Run("test List valida el parametro ACCESS_KEY", func(t *testing.T) {

		_, err := List(mock_list_success)

		if err.Error() != "el ACCESS_KEY es requerido" {
			t.Errorf("Expected: %v, got: %v", "el ACCESS_KEY es requerido", err.Error())
		}
	})

	t.Run("test List error al accesar los servicios externos", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		_, err := List(mock_list_error)

		if err.Error() != "ha ocurrio un error en los servicios api.currencylayer.com" {
			t.Errorf("Expected: %v, got: %v", "ha ocurrio un error en los servicios api.currencylayer.com", err.Error())
		}
	})

	t.Run("test List error al parsear la respuesta del servicio", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		_, err := List(mock_list_success_with_error)

		if err.Error() != "unexpected end of JSON input" {
			t.Errorf("Expected: %v, got: %v", "unexpected end of JSON input", err.Error())
		}
	})

	t.Run("test List error en la respuesta", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		response, _ := List(mock_list_success_with_data_error)

		if response.Success {
			t.Errorf("Expected: %v, got: %v", false, response.Success)
		}
	})

	t.Run("test List success", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		_, err := List(mock_live_success)

		if err != nil {
			t.Errorf("Expected: %v, got: %v", nil, err.Error())
		}
	})
}
