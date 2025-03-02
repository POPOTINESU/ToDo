.PHONY: lint unit-test generate

lint:
	go fmt ./...
	golangci-lint run ./...

unit-test:
	@UNIT_TEST_PACKAGE=$$(go list ./... | grep -v '/tests/' || echo ""); \
	go test -v $$UNIT_TEST_PACKAGE | grep -v 'no test files';

unit-test-with-coverage:
	@UNIT_TEST_PACKAGE=$$(go list ./... | grep -v '/tests/' || echo ""); \
		go test -v -cover -coverprofile=coverage.out $$UNIT_TEST_PACKAGE; \
		go tool cover -html=coverage.out -o unit-coverage-report.html

generate:
	go generate ./...
	git diff --exit-code --name-only | grep '_mock.go' && \
	( echo "Error: Mock files are outdated. Please run 'go generate ./...' and commit the changes." && exit 1 ) || echo "all mocks are latest"