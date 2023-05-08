package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	router := setupServer()

	reqBody := []byte(`{"email": "example@example.com", "password": "example123"}`)
	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

	var responseBody map[string]bool
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Fatal(err)
	}
	if responseBody["login"] != true {
		t.Errorf("expected login to be true but got %v", responseBody["login"])
	}
}

func TestLogout(t *testing.T) {
	router := setupServer()

	req, err := http.NewRequest("DELETE", "/api/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

	var responseBody map[string]bool
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Fatal(err)
	}
	if responseBody["logout"] != true {
		t.Errorf("expected logout to be true but got %v", responseBody["logout"])
	}

	cookies := recorder.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("expected 1 cookie but got %d", len(cookies))
	}
	if cookies[0].Name != "accessToken" {
		t.Errorf("expected cookie name accessToken but got %s", cookies[0].Name)
	}
	if cookies[0].Value != "" {
		t.Errorf("expected cookie value to be empty but got %s", cookies[0].Value)
	}
	if !cookies[0].Expires.IsZero() {
		t.Errorf("expected cookie expiry time to be zero but got %v", cookies[0].Expires)
	}
	if !cookies[0].HttpOnly {
		t.Errorf("expected cookie HttpOnly to be true but got false")
	}
}
