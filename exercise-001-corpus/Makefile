
build:
	go build word_count.go

run:
	./word_count 7oldsamr.txt

dependencies:
	go get -u golang.org/x/lint/golint
	go get -u github.com/stretchr/testify

pre-build:
	gofmt -w .
	golint .

clean:
	rm word_count

.PHONY: build run dependencies clean
