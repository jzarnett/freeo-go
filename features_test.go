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

	expected := `[{"feature":"Plastic Cards","description":"Allow users to have plastic cards"},`
	expected += `{"feature":"External Bookkeeper","description":"Allow inviting an external bookkeeper"},`
	expected += `{"feature":"Teams","description":"Allow assigning users to teams"},`
	expected += `{"feature":"Team Limits","description":"Allow spending lomits at the team level"},`
	expected += `{"feature":"Export to XYZ","description":"Allow export to the XYZ Accounting service"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCompanyFeaturesFree(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/features/d83c34c0-a7a8-42ff-bba6-351d1b647f26", nil)
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

	expected := `[{"feature":"Plastic Cards","access":"upsell"},`
	expected += `{"feature":"External Bookkeeper","access":"upsell"},`
	expected += `{"feature":"Teams","access":"hidden"},`
	expected += `{"feature":"Team Limits","access":"hidden"},`
	expected += `{"feature":"Export to XYZ","access":"visible"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCompanyFeaturesPaid(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/features/3b0628c5-4281-495a-9a5f-789585e95074", nil)
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

	expected := `[{"feature":"Plastic Cards","access":"visible"},`
	expected += `{"feature":"External Bookkeeper","access":"upsell"},`
	expected += `{"feature":"Teams","access":"visible"},`
	expected += `{"feature":"Team Limits","access":"hidden"},`
	expected += `{"feature":"Export to XYZ","access":"visible"}]`
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

func TestPlanFeaturesFree(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/plan/Free/features", nil)
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
	expected := `[{"feature":"Plastic Cards","access":"upsell"},`
	expected += `{"feature":"External Bookkeeper","access":"upsell"},`
	expected += `{"feature":"Teams","access":"hidden"},`
	expected += `{"feature":"Team Limits","access":"hidden"},`
	expected += `{"feature":"Export to XYZ","access":"visible"}]`
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
