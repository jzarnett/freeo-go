package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestAllFeatures(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/features", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	api := router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/features", AllFeatures)
	api.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"feature":"Plastic Cards","description":"Allow users to have plastic cards and not just virtual ones"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCompanyFeatures(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/features/44077d57-38ce-472d-b806-48c87173f76c", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/features/{companyId}", CompanyFeatures)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"feature":"Plastic Cards","access":"hidden"},{"feature":"Self-Onboarding","access":"visible"},{"feature":"Teams","access":"upsell"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCompanyFeaturesUnknown(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/features/unknown", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/features/{companyId}", CompanyFeatures)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	if rr.Body.String() != "" {
		t.Errorf("handler returned unexpected body: got %v want empty",
			rr.Body.String())
	}
}

func TestPlanFeatures(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/plan/basic/features", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/plan/{planName}/features", PlanFeatures)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"feature":"Plastic Cards","access":"upsell"},{"feature":"Self-Onboarding","access":"visible"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPlanFeaturesUnknown(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/plan/unknown/features", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/plan/{planName}/features", PlanFeatures)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	if rr.Body.String() != "" {
		t.Errorf("handler returned unexpected body: got %v want empty",
			rr.Body.String())
	}
}
