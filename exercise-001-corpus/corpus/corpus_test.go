package corpus

import (
	"testing"
)

type testpair struct {
	value  string
	result PairedList
}

var wordCountFromFileTests = []testpair{
	{"test1.txt", PairedList{Pair{Key: "test", Value: 3}}},
}

func TestWordCountFromFile(t *testing.T) {
	for _, pair := range wordCountFromFileTests {
		v, _ := WordCountFromFile(pair.value)
		if !v.Equals(pair.result) {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v)
		}
	}
}

type stringTestPair struct {
	value  string
	result string
}

var stringFromFileTests = []stringTestPair{
	{"test1.txt", "test test test"},
}

func TestStringFromFile(t *testing.T) {
	for _, pair := range stringFromFileTests {
		v, _ := StringFromFile(pair.value)
		if v != pair.result {
			t.Error("For", pair.value,
				"expected", pair.result,
				"got", v)
		}
	}
}

type mapStringTestPair struct {
	result map[string]int
	value  string
}

var wordCountMapFromStringTests = []mapStringTestPair{
	{map[string]int{"test": 3}, "test test test"},
}

func mapContainsAll(a, b map[string]int) bool {
	for k, v := range a {
		if val, ok := b[k]; ok {
			if val != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func TestWordCountMapFromString(t *testing.T) {
	for _, pair := range wordCountMapFromStringTests {
		v := WordCountMapFromString((pair.value))
		if !mapContainsAll(pair.result, v) {
			t.Error("For", pair.value,
				"expected", pair.result,
				"got", v)
		}
	}
}

type mapToPairedListPair struct {
	value  map[string]int
	result PairedList
}

var sortedMapPairedListTests = []mapToPairedListPair{
	{map[string]int{"test": 3, "sample": 2},
		PairedList{Pair{"test", 3}, Pair{"sample", 2}}},
}

func TestSortedMapPairedList(t *testing.T) {
	for _, pair := range sortedMapPairedListTests {
		v := SortedMapPairedList(pair.value, true)
		if !v.Equals(pair.result) {
			t.Error("For", pair.value,
				"expected", pair.result,
				"got", v)
		}
	}
}

var wordCountFromFileSortedTests = []testpair{
	{"test1.txt",
		PairedList{
			Pair{
				Key:   "test",
				Value: 3,
			}}},
}

func TestWordCountFromFileSorted(t *testing.T) {
	for _, pair := range wordCountFromFileSortedTests {
		v, _ := WordCountFromFileSorted(pair.value, true)
		if !v.Equals(pair.result) {
			t.Error("For", pair.value,
				"expected", pair.result,
				"got", v)
		}
	}
}

func BenchmarkWordCountFromFileSorted20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, pair := range wordCountFromFileSortedTests {
			WordCountFromFileSorted(pair.value, true)
		}
	}
}

type pairedListContainsPair struct {
	value    PairedList
	contents Pair
	result   bool
}

var pairedListContainsTests = []pairedListContainsPair{
	{PairedList{Pair{"test", 3}},
		Pair{"test", 3},
		true},
}

func TestPairedList_Contains(t *testing.T) {
	for _, pair := range pairedListContainsTests {
		v := pair.value.Contains(pair.contents)
		if v != pair.result {
			t.Error("For", pair.value,
				"for contents", pair.contents,
				"expected", pair.result,
				"got", v)
		}
	}
}

type pairedListCountPairs struct {
	value   PairedList
	counted Pair
	result  int
}

var pairedListCountTests = []pairedListCountPairs{
	{PairedList{Pair{
		Key:   "test",
		Value: 3,
	},
		Pair{
			Key:   "test",
			Value: 3,
		}},
		Pair{
			Key:   "test",
			Value: 3,
		},
		2},
}

func TestPairedList_Count(t *testing.T) {
	for _, pair := range pairedListCountTests {
		v := pair.value.Count(pair.counted)
		if v != pair.result {
			t.Error("For", pair.value,
				"counted", pair.counted,
				"expected", pair.result,
				"got", v)
		}
	}
}

type pairedListEqualsPair struct {
	lista  PairedList
	listb  PairedList
	result bool
}

var pairedListEqualsPairs = []pairedListEqualsPair{
	{PairedList{Pair{
		Key:   "test",
		Value: 25,
	}},
		PairedList{Pair{
			Key:   "test",
			Value: 25,
		}},
		true},
}

func TestPairedList_Equals(t *testing.T) {
	for _, pair := range pairedListEqualsPairs {
		v := pair.lista.Equals(pair.listb)
		if v != pair.result {
			t.Error("For", pair.lista,
				"and", pair.listb,
				"expected", pair.result,
				"got", v)
		}
	}
}

type pairedListLenPair struct {
	list   PairedList
	length int
}

var pairedListLenTests = []pairedListLenPair{
	{PairedList{Pair{
		Key:   "test",
		Value: 3,
	}, Pair{
		Key:   "sample",
		Value: 1,
	}, Pair{
		Key:   "example",
		Value: 45,
	}}, 3},
}

func TestPairedList_Len(t *testing.T) {
	for _, pair := range pairedListLenTests {
		v := pair.list.Len()
		if v != pair.length {
			t.Error("For", pair.list,
				"expected", pair.length,
				"got", v)
		}
	}
}

type pairedListLessPair struct {
	list   PairedList
	indexa int
	indexb int
	result bool
}

var pairedListLessTests = []pairedListLessPair{
	{PairedList{Pair{
		Key:   "test",
		Value: 3,
	}, Pair{Key: "sample", Value: 2}},
		0,
		1,
		false},
}

func TestPairedList_Less(t *testing.T) {
	for _, pair := range pairedListLessTests {
		v := pair.list.Less(pair.indexa, pair.indexb)
		if v != pair.result {
			t.Error("For", pair.list,
				"index", pair.indexa, "and", pair.indexb,
				"expected", pair.result,
				"got", v)
		}
	}
}
