package main

import (
	"net/http"
	"testing"
)

func TestCompanyFeaturesFree(t *testing.T) {
	res := CompanyFeatures("d83c34c0-a7a8-42ff-bba6-351d1b647f26")

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"feature":"Plastic-Cards","access":"upsell"},`
	expected += `{"feature":"External-Bookkeeper","access":"upsell"},`
	expected += `{"feature":"Teams","access":"hidden"},`
	expected += `{"feature":"Team-Limits","access":"hidden"},`
	expected += `{"feature":"Export-to-XYZ","access":"upsell"}]`
	if res.Body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body, expected)
	}
}

func TestCompanyFeaturesPremium(t *testing.T) {
	res := CompanyFeatures("736af17e-8414-4e31-92f2-64b9e625ac6d")
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"feature":"Plastic-Cards","access":"visible"},`
	expected += `{"feature":"External-Bookkeeper","access":"visible"},`
	expected += `{"feature":"Teams","access":"visible"},`
	expected += `{"feature":"Team-Limits","access":"visible"},`
	expected += `{"feature":"Export-to-XYZ","access":"visible"}]`
	if res.Body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body, expected)
	}
}

func TestCompanyFeaturesEssentials(t *testing.T) {
	res := CompanyFeatures("3b0628c5-4281-495a-9a5f-789585e95074")
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"feature":"Plastic-Cards","access":"visible"},`
	expected += `{"feature":"External-Bookkeeper","access":"visible"},`
	expected += `{"feature":"Teams","access":"upsell"},`
	expected += `{"feature":"Team-Limits","access":"upsell"},`
	expected += `{"feature":"Export-to-XYZ","access":"visible"}]`
	if res.Body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body, expected)
	}
}

func TestCompanyFeaturesUnknown(t *testing.T) {
	res := CompanyFeatures("unknown")
	if status := res.StatusCode; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	if res.Body != "" {
		t.Errorf("handler returned unexpected body: got %v want empty",
			res.Body)
	}
}
