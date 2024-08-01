// Test the main function

package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestMetricsEndpoint(t *testing.T) {
    req, err := http.NewRequest("GET", "/metrics", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := promhttp.Handler()

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("metrics handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check if the Content-Type is correct
    expected := "text/plain; version=0.0.4; charset=utf-8"
    if contentType := rr.Header().Get("Content-Type"); contentType != expected {
        t.Errorf("metrics handler returned unexpected content type: got %v want %v",
            contentType, expected)
    }
}
