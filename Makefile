ci-test: 
	go mod verify\
	&& go mod download\
	&& go test -v --race ./... -coverprofile cover.out\
	&& go tool cover -func cover.out

# Make task for CI golangci-lint report fo CI sonarscanner
ci-golangci-lint-report: 
	go mod verify\
	&& go mod download\
	&& wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.47.1\
	&& ./bin/golangci-lint run --out-format checkstyle --issues-exit-code 0 > golangci-lint.out
	
wire-inject:
	cd internal/factory; \
	wire