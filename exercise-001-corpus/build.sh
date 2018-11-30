gofmt -w .
golint ./...
go clean
go build word_count.go
./word_count 7oldsamr.txt
cd corpus
go test
