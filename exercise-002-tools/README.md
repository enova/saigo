# Go Tools

By installing Go, you now have a basic set of tools that will help you maintain your code format and style. These tools are simple in design, but powerful in function. Here are several that will be invaluable during your time with Go.

## Gofmt

Gofmt simply formats Go programs. Given a file, it operates on that file; given a directory, it operates on all .go files in that directory, recursively.

Apply gofmt's changes to your current directory (and all its subdirectories):

```bash
$ gofmt -w .
```

This command will modify all existing Go files under the current directory. This is really the only version of the command you will probably ever use. Most developers actually have this command wired into their IDE so that whenever they hit save, the code is instantly formatted.

## Golint

Gofmt already takes care of whitespace-related style questions, but it cannot cover things such as variable names. Golint follows the style that is used internally by Google, which is generally accepted by the open source Go community as well.

Golint does not emit errors or warnings, but “suggestions”: These suggestions can be wrong at times, and code that golint complains about isn’t necessarily wrong – it might just be hitting a false positive. Nevertheless, it’s more often right than wrong, and you should definitely run golint on your code from time to time and fix those suggestions that it is right about.

Install and run golint over your current directory:

```bash
$ go get github.com/golang/lint/golint
$ golint ./...
```

Fix any issues raised by golint. You can read more about it at https://github.com/golang/lint.

## Go vet

Go vet is concerned with correctness, whereas golint is concerned with coding style. You can read more about it at https://golang.org/cmd/vet/.

Go comes with its own linter, go vet, for analysing Go code and finding common mistakes. So far, it can check your code for the following mistakes:

* Useless assignments
* Common mistakes when using sync/atomic
* Invalid +build tags
* Using composite literals without keyed fields
* Passing locks by value
* Comparing functions with nil
* Using wrong printf format specifiers
* Closing over loop variables the wrong way
* Struct tags that do not follow the canonical format
* Unreachable code
* Misuse of unsafe.Pointer
* Mistakes involving boolean operators

Run go vet over your current directory:

```bash
$ go vet ./...
```

Fix any issues raised by go vet.

## Cover

Code coverage reports provide a quick and easy way of finding untested code. Go’s support for these reports consists of two components: Support in go test for generating coverage profiles and go tool cover to generate reports from it.

Set the -coverprofile option while running `go test` and it will generate a coverage profile:

```bash
go test ./src/words -coverprofile=coverage.out
```

You can then pass the profile to `go tool cover` and generate different kinds of reports. If you want per-function coverage, you can use the `-func` flag.

```bash
go tool cover -func=coverage.out
```

If you want a HTML presentation of the source code, you can use the `-html` flag:
```bash
go tool cover -html=coverage.out
```

To view code coverage analysis for each build of your source code, check out [Coveralls](https://coveralls.io/).

### Extra

Several other tools are listed at
http://dominik.honnef.co/posts/2014/12/an_incomplete_list_of_go_tools/ and https://golang.org/cmd/. Read through them to learn about what else can be in your toolset.
