Package main 

import (
	"testing"
)


func word_count_test(t *testing.T) {
	words := scanWords("../7oldsamr.txt")

	length := len(words)
	v1 := words[0]
	v2 := words[100]

	if length != 497 {
		t.Error(
			"For the length, expected 497, got", length,
		)
	}
	if v1 != "THE" {
		t.Error(
			"For words[0], expected THE, got", words[0],
		)
	}
	if v2 != "Emperor" {
		t.Error(
			"For words[100], expected Emperor, got", words[100],
		)
	}
}
