## Description
Write a console application that counts the frequencies of words that appear in a text document.

## Input
Your application (say it's called `word_count`) should accept a single command-line argument containing the path of the text file.

Example:
```
$ word_count 7oldsamr.txt
```

## Output
Your application should produce a list of words and their frequencies _sorted in reverse numerical order_.

Example:
```
36 the
18 and
14 to
12 a
10 old
10 of
10 in
9 they
7 for
6 you
6 seven
6 samurai
6 robbers
6 raiko
...
```

In case of ties, any ordering is fine as long as it is deterministic (so that you can write proper tests!) To simplify things you may consider uppercase letters indistinguishable from lowercase letters.

## Implementation
The purpose of this exercise is to teach usage of basic Go elements and packages like `map, sort, strings`. So please refrain from installing custom packages. Your imports should look something like this:

```
import (
  "fmt"
  "sort"
  "strings"
)
```
and not like this:
```
import (
  "github.com/hotshot/nlp"
  "github.com/topcoder1/histogram"
)
```
## (Recommended) Directory Structure

Within your project directory there should be three files, two under a subdirectory named `corpus` (this is the package name):

```
corpus/corpus.go
corpus/corpus_test.go
word_count.go
```

The first two files within the `corpus/` directory are part of the `corpus` package you are building. Every Go file in that directory should start with `package corpus`. The filenames themselves can be arbitrary. Test-Code should reside in filenames ending in `_test.go`. The third file `word_count.go` will contain your executable code and should begin with `package main`.

You will probably want to run the standard Go formatter and linter over your code whenever you build. You can add a script `build.sh` in the project directory that looks something like this:

```
build.sh
--------

gofmt -w .
golint ./...
go build word_count.go
```
Make sure the script is executable:
```
$ chmod u+x build.sh
```
Then any time you want to build your code just run that script (`$ ./build.sh`). The first line `gofmt -w .` will format your code and overwrite all existing Go files in your project directory. The second line `golint ./...` lets you know which style-conventions are not being observed, but it will not overwrite your code. Each of these steps helps make all Go code look similar. Almost all Go developers use them in their work.

The third line `go build word_count.go` should create a (binary) executable `word_count`. You can then run your executable from the console `./word_count 7oldsamr.txt`.

## Testing
Please include a test of your software that runs with `go test ./...` The [testify/assert](http://github.com/stretchr/testify) package is a testing package that is widely used by the Go community, and we recommend using it for all Saigo exercises. For example:

```go
corpus_test.go
--------------

package corpus

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCorpus(t *testing.T) {
  result := Analyze("Are you serious? I knew you were.")
  assert.Equal(t, len(result), 5)
  assert.Equal(t, result[0].Word, "you")
  assert.Equal(t, result[0].Count, 2)
}
```

## Benchmarking (Optional)
Try to use the Go [benchmarking](http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) tools that run with `go test -bench=.` 
