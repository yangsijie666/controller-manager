.PHONY: generate
# generate
generate:
	go mod tidy
	go generate ./...
