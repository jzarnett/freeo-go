package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

const (
	HIDDEN  string = "hidden"
	VISIBLE string = "visible"
	UPSELL  string = "upsell"
)

type featureAccess struct {
	id          int
	Feature     string `json:"feature"`
	Access      string `json:"access,omitempty"`
	Description string `json:"description,omitempty"`
}

func getAllFeatures() []featureAccess {
	// Dummy implementation
	f1 := featureAccess{
		Feature:     "Plastic Cards",
		Description: "Allow users to have plastic cards"}
	f2 := featureAccess{
		Feature:     "External Bookkeeper",
		Description: "Allow inviting an external bookkeeper"}
	f3 := featureAccess{
		Feature:     "Teams",
		Description: "Allow assigning users to teams"}
	f4 := featureAccess{
		Feature:     "Team Limits",
		Description: "Allow spending lomits at the team level"}
	f5 := featureAccess{
		Feature:     "Export to XYZ",
		Description: "Allow export to the XYZ Accounting service"}
	features := make([]featureAccess, 5)
	features[0] = f1
	features[1] = f2
	features[2] = f3
	features[3] = f4
	features[4] = f5
	return features
}

func getFeaturesForCompany(companyID string) []featureAccess {
	// Dummy implementation
	if companyID == "d83c34c0-a7a8-42ff-bba6-351d1b647f26" {
		// FREE tier company
		f1 := featureAccess{
			Feature: "Plastic Cards",
			Access:  UPSELL}
		f2 := featureAccess{
			Feature: "External Bookkeeper",
			Access:  UPSELL}
		f3 := featureAccess{
			Feature: "Teams",
			Access:  HIDDEN}
		f4 := featureAccess{
			Feature: "Team Limits",
			Access:  HIDDEN}
		f5 := featureAccess{
			Feature: "Export to XYZ",
			Access:  VISIBLE}
		features := make([]featureAccess, 5)
		features[0] = f1
		features[1] = f2
		features[2] = f3
		features[3] = f4
		features[4] = f5
		return features
	} else if companyID == "3b0628c5-4281-495a-9a5f-789585e95074" {
		f1 := featureAccess{
			Feature: "Plastic Cards",
			Access:  VISIBLE}
		f2 := featureAccess{
			Feature: "External Bookkeeper",
			Access:  VISIBLE}
		f3 := featureAccess{
			Feature: "Teams",
			Access:  VISIBLE}
		f4 := featureAccess{
			Feature: "Team Limits",
			Access:  VISIBLE}
		f5 := featureAccess{
			Feature: "Export to XYZ",
			Access:  VISIBLE}
		features := make([]featureAccess, 5)
		features[0] = f1
		features[1] = f2
		features[2] = f3
		features[3] = f4
		features[4] = f5
		return features
	}
	return nil
}

func getFeaturesForPlan(plan string) []featureAccess {
	// Dummy implementation
	if plan == "Free" {
		// FREE tier company
		f1 := featureAccess{
			Feature: "Plastic Cards",
			Access:  UPSELL}
		f2 := featureAccess{
			Feature: "External Bookkeeper",
			Access:  UPSELL}
		f3 := featureAccess{
			Feature: "Teams",
			Access:  HIDDEN}
		f4 := featureAccess{
			Feature: "Team Limits",
			Access:  HIDDEN}
		f5 := featureAccess{
			Feature: "Export to XYZ",
			Access:  VISIBLE}
		features := make([]featureAccess, 5)
		features[0] = f1
		features[1] = f2
		features[2] = f3
		features[3] = f4
		features[4] = f5
		return features
	} else if plan == "Essential" || plan == "Premium" {
		f1 := featureAccess{
			Feature: "Plastic Cards",
			Access:  VISIBLE}
		f2 := featureAccess{
			Feature: "External Bookkeeper",
			Access:  UPSELL}
		f3 := featureAccess{
			Feature: "Teams",
			Access:  VISIBLE}
		f4 := featureAccess{
			Feature: "Team Limits",
			Access:  HIDDEN}
		f5 := featureAccess{
			Feature: "Export to XYZ",
			Access:  VISIBLE}
		features := make([]featureAccess, 5)
		features[0] = f1
		features[1] = f2
		features[2] = f3
		features[3] = f4
		features[4] = f5
		return features
	} else if plan == "Pro" {
		f1 := featureAccess{
			Feature: "Plastic Cards",
			Access:  VISIBLE}
		f2 := featureAccess{
			Feature: "External Bookkeeper",
			Access:  VISIBLE}
		f3 := featureAccess{
			Feature: "Teams",
			Access:  VISIBLE}
		f4 := featureAccess{
			Feature: "Team Limits",
			Access:  VISIBLE}
		f5 := featureAccess{
			Feature: "Export to XYZ",
			Access:  VISIBLE}
		features := make([]featureAccess, 5)
		features[0] = f1
		features[1] = f2
		features[2] = f3
		features[3] = f4
		features[4] = f5
		return features
	}
	return nil
}

func buildResponse(features []featureAccess) events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	if features == nil {
		resp.StatusCode = 404
		return resp
	}

	marshalled, err := json.Marshal(features)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = err.Error()
		return resp
	}

	resp.StatusCode = 200
	resp.Headers["Content-Type"] = "application/json"
	resp.Body = string(marshalled[:])
	return resp
}

// AllFeatures - Handler for requesting all features
func AllFeatures() events.APIGatewayProxyResponse {
	features := getAllFeatures()
	return buildResponse(features)
}

// CompanyFeatures - handler for finding what features this company has
func CompanyFeatures(companyID string) events.APIGatewayProxyResponse {
	features := getFeaturesForCompany(companyID)
	return buildResponse(features)
}

// PlanFeatures - handler for finding what features are in a given plan
func PlanFeatures(plan string) events.APIGatewayProxyResponse {
	features := getFeaturesForPlan(plan)
	return buildResponse(features)
}

/*
func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/features", AllFeatures).Methods(http.MethodGet)
	api.HandleFunc("/features/{companyId}", CompanyFeatures).Methods(http.MethodGet)
	api.HandleFunc("/plan/{planName}/features", PlanFeatures).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
*/
