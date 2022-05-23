echo "** running formatters"

gofmt -s -w $(find . -type f -name '*.go' -not -path "./vendor/*")

echo "** running golangci-lint"
golangci-lint run ./...

echo "** running tests"
go test $(go list ./... | grep -v mocks) -cover
