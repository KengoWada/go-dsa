COVERAGE_HTML = ./index.out.html
COVERAGE_FILE = ./.coverage.out

test:
	@go test -v ./...

test-cov:
	@go test -v ./... -coverprofile=$(COVERAGE_FILE) && go tool cover -html=$(COVERAGE_FILE) -o=$(COVERAGE_HTML)

.PHONY: test test-cov