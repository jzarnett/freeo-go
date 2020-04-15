#! /bin/bash

mkdir -p build

env GOOS=linux GOARCH=amd64 go build -o build/allfeatures src/allfeatures.go src/common.go
zip -j build/allfeatures.zip build/allfeatures

env GOOS=linux GOARCH=amd64 go build -o build/companyfeatures src/companyfeatures.go src/common.go
zip -j build/companyfeatures.zip build/companyfeatures

env GOOS=linux GOARCH=amd64 go build -o build/planfeatures src/planfeatures.go src/common.go
zip -j build/planfeatures.zip build/planfeatures
