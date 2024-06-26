package handlers_test

import (
	"io"
	"net/http"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestHealthCheck(t *testing.T) {
	baseUrl := "http://localhost:3000"

	res, err := http.Get(baseUrl + "/api/healthcheck")

	if err != nil {
		t.Errorf("Error %s", err.Error())
		return
	}

	if http.StatusOK != res.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error parsing response body: %s", err.Error())
		return
	}

	bodyString := string(bodyBytes)

	if bodyString != "HealthCheck OK!" {
		t.Errorf("Expected an empty array. Got %s", bodyString)
	}
}
