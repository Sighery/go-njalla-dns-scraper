installdeps:
	go get
test:
	go test ./...
build:
	go build -o njallaclient

.PHONY: installdeps
.PHONY: test
