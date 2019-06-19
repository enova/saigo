package main

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	dict, err := ScanWords("../7oldsamr.txt")
	if err != nil {
		panic(err)
	}
	v1 := dict["the"]
	v2 := dict["and"]
	v3 := dict["to"]
	if v1 != 36 {
		t.Error(
			"For 'the' expected 36, got", v1,
		)
	}
	if v2 != 18 {
		t.Error(
			"For 'and' expected 18, got", v2,
		)
	}
	if v3 != 14 {
		t.Error(
			"For 'to' expected 14, got", v3,
		)
	}
}
