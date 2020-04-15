package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func allFeaturesHandler(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return AllFeatures()
}

func main() {
	lambda.Start(allFeaturesHandler)
}
