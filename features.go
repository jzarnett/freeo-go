package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	HIDDEN  string = "hidden"
	VISIBLE string = "visible"
	UPSELL  string = "upsell"
)

type featureAccess struct {
	Feature     string `json:"feature"`
	Access      string `json:"access,omitempty"`
	Description string `json:"description,omitempty"`
}

func getAllFeatures() []featureAccess {
	// Dummy implementation
	f := featureAccess{
		Feature:     "Plastic Cards",
		Description: "Allow users to have plastic cards and not just virtual ones"}
	features := make([]featureAccess, 1)
	features[0] = f
	return features
}

func getFeaturesForCompany(companyID string) []featureAccess {
	// Dummy implementation
	if companyID == "unknown" {
		return nil
	}

	f1 := featureAccess{
		Feature: "Plastic Cards",
		Access:  HIDDEN}
	f2 := featureAccess{
		Feature: "Self-Onboarding",
		Access:  VISIBLE}
	f3 := featureAccess{
		Feature: "Teams",
		Access:  UPSELL}
	features := make([]featureAccess, 3)
	features[0] = f1
	features[1] = f2
	features[2] = f3

	return features
}

func getFeaturesForPlan(plan string) []featureAccess {
	// Dummy implementation
	if plan == "unknown" {
		return nil
	}

	f1 := featureAccess{
		Feature: "Plastic Cards",
		Access:  UPSELL}
	f2 := featureAccess{
		Feature: "Self-Onboarding",
		Access:  VISIBLE}
	features := make([]featureAccess, 2)
	features[0] = f1
	features[1] = f2

	return features
}

func writeResult(w http.ResponseWriter, features []featureAccess) {
	if features == nil {
		w.WriteHeader(404)
		return
	}

	marshalled, err := json.Marshal(features)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(marshalled)
}

// AllFeatures - Handler for requesting all features
func AllFeatures(w http.ResponseWriter, r *http.Request) {
	features := getAllFeatures()
	writeResult(w, features)
}

// CompanyFeatures - handler for finding what features this company has
func CompanyFeatures(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	features := getFeaturesForCompany(pathParams["companyId"])
	writeResult(w, features)
}

// PlanFeatures - handler for finding what features are in a given plan
func PlanFeatures(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	features := getFeaturesForPlan(pathParams["planName"])
	writeResult(w, features)
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/features", AllFeatures).Methods(http.MethodGet)
	api.HandleFunc("/features/{companyId}", CompanyFeatures).Methods(http.MethodGet)
	api.HandleFunc("/plan/{planName}/features", PlanFeatures).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
