package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func planFeaturesHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return PlanFeatures(req.PathParameters["planName"]), nil
}

func main() {
	lambda.Start(planFeaturesHandler)
}
