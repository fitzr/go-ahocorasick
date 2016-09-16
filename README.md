go ahocorasick
===========
[![Go Report Card](https://goreportcard.com/badge/github.com/fitzr/goahocorasick)](https://goreportcard.com/report/github.com/fitzr/goahocorasick) [![GoDoc](https://godoc.org/github.com/fitzr/goahocorasick?status.png)](https://godoc.org/github.com/fitzr/goahocorasick) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/fitzr/goahocorasick/blob/master/LICENSE)

Aho-Corasick algorithm in golang.

~~~ go
package main

import (
    "fmt"
    "github.com/fitzr/goahocorasick"
)

func main() {

    keywords := []string{"key", "word", "keyword", "ーワー"}
    a := goahocorasick.New(keywords)

    target := "keyword is キーワード"
    results := a.Match(target)

    for _, result := range results {
        fmt.Println(result)
    }
}
~~~

~~~
# return [index length] in units of utf8
[0 3]  (key)
[0 7]  (keyword)
[3 4]  (word)
[12 3] (ーワー)
~~~

[大規模サービス技術入門](http://gihyo.jp/book/2010/978-4-7741-4307-1) Lesson 22-23 課題実装
