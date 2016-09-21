package goahocorasick_test

import (
	. "github.com/fitzr/goahocorasick"
	"reflect"
	"testing"
)

func TestAhoCorasick_Match(t *testing.T) {
	input := "a his hoge hershe xx."
	keywords := []string{"he", "hers", "his", "she"}
	expected := [][]int{
		{2, 3},  // his
		{11, 2}, // he
		{11, 4}, // hers
		{14, 3}, // she
		{15, 2}} // he
	expectMatched := []string{"his", "he", "hers", "she", "he"}

	sut := New(keywords)
	actual := sut.Match(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}

	matched := make([]string, len(actual))
	for i, v := range actual {
		matched[i] = input[v[0] : v[0]+v[1]]
	}
	if !reflect.DeepEqual(expectMatched, matched) {
		t.Errorf("\nexpected: %v\nactual: %v", expectMatched, matched)
	}
}

func TestAhoCorasick_MatchWithMiddleWord(t *testing.T) {
	input := "sheeg"
	keywords := []string{"shees", "he"}
	expected := [][]int{{1, 2}} // he

	sut := New(keywords)
	actual := sut.Match(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}
