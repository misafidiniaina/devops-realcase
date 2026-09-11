package server

import (
	"net/http/httptest"
	"testing"
)

func TestHealthWithoutDatabase(t *testing.T) {
	r := httptest.NewRecorder()
	New(nil).health(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 200 {
		t.Fatalf("got %d", r.Code)
	}
	if r.Body.String() != "{\"status\":\"ok\"}\n" {
		t.Fatalf("unexpected body: %s", r.Body)
	}
}
func TestMetrics(t *testing.T) {
	r := httptest.NewRecorder()
	New(nil).metrics(r, httptest.NewRequest("GET", "/metrics", nil))
	if r.Code != 200 || r.Body.Len() == 0 {
		t.Fatal("metrics endpoint did not respond")
	}
}
