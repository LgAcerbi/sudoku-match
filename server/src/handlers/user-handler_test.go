package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

// TO-DO adjust expect
func TestRegisterUser(t *testing.T) {
	baseUrl := "http://localhost:3000"

	body, _ := json.Marshal(map[string]string{
		"email":         "test@mail.com",
		"googleIdToken": "some-token",
		"nickname":      "some-nick-name",
	})
	payload := bytes.NewBuffer(body)

	res, err := http.Post(baseUrl+"/api/register", "application/json", payload)

	if err != nil {
		t.Errorf("Error %s", err.Error())
		return
	}

	if http.StatusCreated != res.StatusCode {
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

// TO-DO adjust expect
func TestFindUserByEmail(t *testing.T) {
	baseUrl := "http://localhost:3000"

	email := "test@mail.com"

	res, err := http.Get(baseUrl + "/api/register" + email)

	if err != nil {
		t.Errorf("Error %s", err.Error())
		return
	}

	if http.StatusCreated != res.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error parsing response body: %s", err.Error())
		return
	}

	bodyString := string(bodyBytes)

	if bodyString != "" {
		t.Errorf("Expected an empty array. Got %s", bodyString)
	}
}
