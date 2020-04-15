# Freeo Features Tool

This is a pretty simple webservice intended to serve as a way to quickly and easily
look up feature details within Pleo. This is a supporting tool for the 2020-04 
Hackathon in the Freeo team. It

## Endpoints

1. `/v1/features` - returns a list of all features, including their name and a description.
2. `/v1/features/{companyId}` - returns a list of the features for this company, including their name and their visibility (visible, hidden, or upsell).
3. `/v1/{planName}/features` - returns a list of the features associated with a specific plan, including their name and their visibility  (visible, hidden, or upsell).

## Requirements

`go get -u github.com/gorilla/mux`
`go get github.com/aws/aws-lambda-go/lambda`

## Building

`./build.sh`

## Testing

`./test.sh`