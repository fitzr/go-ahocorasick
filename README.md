ahocorasick
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
# [index length]
[0 3]  (key)
[0 7]  (keyword)
[3 4]  (word)
[12 3] (ーワー)
~~~

#### License
MIT