.PHONY: golang-onboarding
golang-onboarding: main.go
	@go build -o golang-onboarding main.go

.PHONY: download
download:
	@GOSUMDB=off go mod download -x
	@go mod vendor

.PHONY: mocks
mocks:
	@which mockgen 2>/dev/null || echo "please install github.com/golang/mock"
	@go generate $$(go list ./... | grep -v vendor)
	@echo "=== Mocks Regenerated ==="
