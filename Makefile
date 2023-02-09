all: server-darwin-arm64 server-amd64 server-arm64 client-darwin-arm64 client-amd64 client-arm64

server-darwin-arm64: server.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $@ $<

server-amd64: server.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ $<

server-arm64: server.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $@ $<

client-darwin-arm64: client.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $@ $<

client-amd64: client.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ $<

client-arm64: client.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $@ $<
