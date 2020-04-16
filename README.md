# Freeo Features Tool

This is a pretty simple webservice intended to serve as a way to quickly and easily
look up feature details within Pleo. This is a supporting tool for the 2020-04 
Hackathon in the Freeo team.

## Functionality

This code creates three AWS lambdas:

1. `allFeatures` - returns all features' names and descriptions.
2. `companyFeatures` - returns the feature names that a company has access to in its plan. The company UUID must be provided as a path parameter.
3. `planFeatures` - returns the features that a particular plan contains. The plan name (e.g., "Free") must be provided as a path parameter.

All functions return a JSON response.

A feature, when part of a plan or company's feature set, will also include the access level, which will be one of: {hidden,upsell,visible}. Upsell features are not available in the plan but they are shown to the user as an incentive to get them to upgrade to a higher plan. 

A sample JSON response describing a feature overall is:
`{"feature":"Plastic-Cards","description":"Allow users to have plastic cards"},`

A sample JSON response describing a feature's access is:
`{"feature":"Plastic-Cards","access":"upsell"}`

## Requirements

`go get github.com/aws/aws-lambda-go/lambda`

## Building

`./build.sh` produces 3 ZIP files that are ready for uploading into AWS.

## Testing

`./test.sh`