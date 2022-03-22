package registry

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceDiscoveryHandler(t *testing.T) {
	// Create Request to test service discovery protocol
	req, err := http.NewRequest("GET", "/.well-known/terraform.json", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create Handler and Recorder for the test
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ServiceDiscoveryHandler)

	handler.ServeHTTP(rr, req)

	// Check the Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"modules.v1":"/terraform/modules/v1/"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
