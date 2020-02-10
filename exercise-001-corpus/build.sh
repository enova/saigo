gofmt -w .
golint ./...
go build word_count.go
pushd corpus
go test
go test -bench=.
popd
./word_count -file 7oldsamr.txt
