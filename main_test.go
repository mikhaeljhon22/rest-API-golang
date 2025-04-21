package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func TestAllArticles(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Buat response recorder (penangkap output handler)
	rr := httptest.NewRecorder()

	// Assign handler yang mau diuji
	handler := http.HandlerFunc(allArticles)

	// Jalankan handler dengan request palsu
	handler.ServeHTTP(rr, req)

	// Cek status code apakah 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Expected hasilnya
	expected := Articles{
		{Title: "Test title", Desc: "Test desc", Content: "Hello World"},
	}

	// Decode hasil respons JSON
	var got Articles
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Errorf("could not parse response body: %v", err)
	}

	// Bandingkan hasil dengan expected
	if len(got) != 1 || got[0] != expected[0] {
		t.Errorf("handler returned unexpected body: got %v want %v", got, expected)
	}
}
