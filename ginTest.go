package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserStruct(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/get/struct", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var user User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	if err != nil {
		t.Fatalf("Error unmarshalling response: %v", err)
	}

	if user.Username != "Mikhael" || user.Password != "Akuanakkaya123" {
		t.Errorf("Unexpected user data: %+v", user)
	}
}
