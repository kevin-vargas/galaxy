go test -coverprofile=coverage.out $(go list ./... | grep -v mocks)
go tool cover -html=coverage.out
