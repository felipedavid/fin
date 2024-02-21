
build:
	@templ generate
	@go build -o ./tmp/main.exe cmd/main.go

fmt:
	@templ fmt .
	@go fmt .\...