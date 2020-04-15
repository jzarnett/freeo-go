package main

import (
	"net/http"
	"testing"
)

func TestAllFeatures(t *testing.T) {
	res := AllFeatures()
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"feature":"Plastic Cards","description":"Allow users to have plastic cards"},`
	expected += `{"feature":"External Bookkeeper","description":"Allow inviting an external bookkeeper"},`
	expected += `{"feature":"Teams","description":"Allow assigning users to teams"},`
	expected += `{"feature":"Team Limits","description":"Allow spending lomits at the team level"},`
	expected += `{"feature":"Export to XYZ","description":"Allow export to the XYZ Accounting service"}]`
	if res.Body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body, expected)
	}
}
