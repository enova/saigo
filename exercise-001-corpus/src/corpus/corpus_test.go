package corpus

import (
    "testing"
    "fmt"
)

func TestWordCount(t *testing.T) {
	fmt.Printf("Running first test..\n")
	WordCount("../../7oldsamr.txt")
}
