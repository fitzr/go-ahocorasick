go ahocorasick
===========
[![Go Report Card](https://goreportcard.com/badge/github.com/fitzr/goahocorasick)](https://goreportcard.com/report/github.com/fitzr/goahocorasick)

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

[godoc](https://godoc.org/github.com/fitzr/goahocorasick)

[大規模サービス技術入門](http://gihyo.jp/book/2010/978-4-7741-4307-1) Lesson 22-23 課題実装

#### License
MIT