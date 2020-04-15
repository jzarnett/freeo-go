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

var everyFeature = []featureAccess{
	featureAccess{
		Feature:     "Plastic-Cards",
		Description: "Allow users to have plastic cards"},
	featureAccess{
		Feature:     "External-Bookkeeper",
		Description: "Allow inviting an external bookkeeper"},
	featureAccess{
		Feature:     "Teams",
		Description: "Allow assigning users to teams"},
	featureAccess{
		Feature:     "Team-Limits",
		Description: "Allow spending limits at the team level"},
	featureAccess{
		Feature:     "Export-to-XYZ",
		Description: "Allow export to the XYZ Accounting service"}}

var premiumCompanyUUID = "736af17e-8414-4e31-92f2-64b9e625ac6d"
var freeTierCompanyUUID = "d83c34c0-a7a8-42ff-bba6-351d1b647f26"
var otherCompanyUUID = "3b0628c5-4281-495a-9a5f-789585e95074"

var planMap = map[string][]featureAccess{
	"Free": {featureAccess{Feature: "Plastic-Cards", Access: UPSELL},
		featureAccess{Feature: "External-Bookkeeper", Access: UPSELL},
		featureAccess{Feature: "Teams", Access: HIDDEN},
		featureAccess{Feature: "Team-Limits", Access: HIDDEN},
		featureAccess{Feature: "Export-to-XYZ", Access: UPSELL}},
	"Essential": {featureAccess{Feature: "Plastic-Cards", Access: VISIBLE},
		featureAccess{Feature: "External-Bookkeeper", Access: VISIBLE},
		featureAccess{Feature: "Teams", Access: UPSELL},
		featureAccess{Feature: "Team-Limits", Access: UPSELL},
		featureAccess{Feature: "Export-to-XYZ", Access: VISIBLE}},
	"Premium": {featureAccess{Feature: "Plastic-Cards", Access: VISIBLE},
		featureAccess{Feature: "External-Bookkeeper", Access: VISIBLE},
		featureAccess{Feature: "Teams", Access: VISIBLE},
		featureAccess{Feature: "Team-Limits", Access: VISIBLE},
		featureAccess{Feature: "Export-to-XYZ", Access: VISIBLE}},
	"Pro": {featureAccess{Feature: "Plastic-Cards", Access: VISIBLE},
		featureAccess{Feature: "External-Bookkeeper", Access: VISIBLE},
		featureAccess{Feature: "Teams", Access: VISIBLE},
		featureAccess{Feature: "Team-Limits", Access: VISIBLE},
		featureAccess{Feature: "Export-to-XYZ", Access: VISIBLE}}}

var companyMap = map[string][]featureAccess{
	premiumCompanyUUID:  planMap["Premium"],
	freeTierCompanyUUID: planMap["Free"],
	otherCompanyUUID:    planMap["Essential"]}

func getAllFeatures() []featureAccess {
	// Dummy implementation
	return everyFeature
}

func getFeaturesForCompany(companyID string) []featureAccess {
	// Dummy implementation
	return companyMap[companyID]
}

func getFeaturesForPlan(plan string) []featureAccess {
	return planMap[plan]
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
	resp.Headers["Access-Control-Allow-Origin"] = "*"
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
