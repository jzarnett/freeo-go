package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func allFeaturesHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return AllFeatures(), nil
}

func main() {
	lambda.Start(allFeaturesHandler)
}
