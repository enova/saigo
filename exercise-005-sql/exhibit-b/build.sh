gofmt -w .
golint ./...
go build db.go
