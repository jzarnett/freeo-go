package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func companyFeaturesHandler(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return CompanyFeatures(req.PathParameters["companyId"])
}

func main() {
	lambda.Start(companyFeaturesHandler)
}
