#! /bin/bash

go test -v src/allfeatures.go src/common.go src/allfeatures_test.go 
go test -v src/companyfeatures.go src/common.go src/companyfeatures_test.go 
go test -v src/planfeatures.go src/common.go src/planfeatures_test.go 