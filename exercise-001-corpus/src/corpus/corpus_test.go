package corpus

import (
	"fmt"
	"testing"
)

func TestWordCount(t *testing.T) {
	fmt.Printf("Running first test..\n")
	WordCount("../../7oldsamr.txt")
}
