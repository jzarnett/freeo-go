package main

import (
	"net/http"
	"testing"
)

func TestPlanFeaturesFree(t *testing.T) {
	res := PlanFeatures("Free")
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

func TestPlanFeaturesUnknown(t *testing.T) {
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
