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

## Testing
Please include a test of your software that runs with `go test ./...`

## Benchmarking (Optional)
Try to use the Go [benchmarking](http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) tools that run with `go test -bench=.` 
