package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func planFeaturesHandler(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return PlanFeatures(req.PathParameters["planName"])
}

func main() {
	lambda.Start(planFeaturesHandler)
}