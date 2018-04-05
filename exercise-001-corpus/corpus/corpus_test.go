package corpus

import (
	"reflect"
	"testing"
)

var result string

type testTypes struct {
	testStringFromFile      string
	testWordArray           []string
	testWordMap             wordCountList
	testCleanedFileContents string
	testOutput              string
}

var testData = testTypes{
	testStringFromFile: `Test string "this"
                        string is the best string.`,
	testCleanedFileContents: "Test string this string is the best string",
	testWordArray:           []string{"Test", "this", "string", "this", "string", "string"},
	testWordMap: []wordCountPair{
		wordCountPair{"string", 3},
		wordCountPair{"this", 2},
		wordCountPair{"Test", 1},
	},
	testOutput: "\n2 testing\n1 Testing\n1 Anyone\n1 there",
}

func TestCleanFileContents(t *testing.T) {
	contents, err := cleanFileContents(testData.testStringFromFile)

	if contents != testData.testCleanedFileContents && err == nil {
		t.Error(
			"\nFor: ", testData.testStringFromFile,
			"\nExpected: ", testData.testCleanedFileContents,
			"\nGot: ", contents,
		)
	}
}

func TestCountWords(t *testing.T) {
	testMap := countWords(testData.testWordArray)
	if !reflect.DeepEqual(testMap, testData.testWordMap) {
		t.Error(
			"\nFor: ", testData.testWordArray,
			"\nExpected: ", testData.testWordMap,
			"\nGot: ", testMap,
		)
	}
}

func TestParseFile(t *testing.T) {
	parseResult := ParseFile("test_file.txt")
	if len(parseResult) != len(testData.testOutput) {
		t.Error(
			"\nFor: ", "test_file.txt",
			"\nExpected: ", testData.testOutput,
			"\nGot: ", parseResult)
	}
}

func BenchmarkParseFile(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = ParseFile("test_file.txt")
	}
	result = r
}
