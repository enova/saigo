
BIN_NAME := out

$(BIN_NAME): *.go
	go build -o $@ .

run:
	./$(BIN_NAME) 7oldsamr.txt

dependencies:
	go get -u golang.org/x/lint/golint
	go get -u github.com/stretchr/testify

pre-build:
	gofmt -w .
	golint .

clean:
	rm $(BIN_NAME)

.PHONY: run dependencies clean
