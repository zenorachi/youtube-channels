.PHONY: swagger
swagger:
	swag fmt
	swag init -g ./internal/app/app.go

.PHONY: lint
lint:
	$(GOPATH)/bin/golangci-lint run -c .golangci.yml

.PHONY: imports
imports:
	$(GOPATH)/bin/goimports -w ./cmd ./internal/ ./pkg
