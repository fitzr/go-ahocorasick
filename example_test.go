package goahocorasick_test

import (
	"fmt"
	"github.com/fitzr/goahocorasick"
)

func ExampleAhocorasick() {
	keywords := []string{"key", "word", "keyword", "ーワー"}
	a := goahocorasick.New(keywords)

	target := "keyword is キーワード"
	results := a.Match(target)

	for _, result := range results {
		fmt.Println(result)
	}

	// Output:
	// [0 3]
	// [0 7]
	// [3 4]
	// [12 3]
}
