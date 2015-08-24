BINARY	= bin/uptime-httpd
GO	= env GOPATH="$(PWD)/vendor:$(PWD)" go

all:
	@echo "make release   # Build a release version"
	@echo "make run       # Run a development version"

release: clean
	@env GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY) uptime-httpd.go

run: clean
	@$(GO) build -o $(BINARY) uptime-httpd.go
	@./$(BINARY)

clean:
	rm -fr bin pkg vendor/pkg