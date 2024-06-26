MAIN_PKG := ./pkg/hexe

.PHONY: no-dirty
no-dirty:
	git diff --exit-code

.PHONY: generate
generate:
	go generate ./...
	go fmt ./...

.PHONY: tidy
tidy: generate
	go mod tidy -v

.PHONY: audit
audit: tidy no-dirty
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -buildvcs -vet=off ./...

.PHONY: test
test:
	go test -v -buildvcs ./...

.PHONY: coverage
coverage:
	go test -v -buildvcs -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: examples
examples:
	go run ./examples
