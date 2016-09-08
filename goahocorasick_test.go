package goahocorasick

import (
	"bufio"
	"os"
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

func TestAhoCorasick_MatchWithBigData(t *testing.T) {
	input := "今日は天気がよかったので、近くの海まで愛犬のしなもんと一緒にお散歩。写真は海辺を楽しそうに歩くしなもん。そのあとついでにお買い物にも行ってきました。「はてなの本」を買ったので、はてなダイアリーの便利な商品紹介ツール「はまぞう」を使って紹介してみるよ。とてもおもしろいのでみんなも読んでみてね。"
	keywords := getKeywords()
	expectMatched := []string{"今日", "天気", "しなもん", "散歩", "写真", "海辺", "しなもん", "はてな", "はてなの本", "はてな", "はてなダイアリ", "はてなダイアリー", "ダイアリー", "商品", "はまぞう", "おもしろい"}

	sut := New(keywords)
	actual := sut.Match(input)

	matched := make([]string, len(actual))
	runes := []rune(input)
	for i, v := range actual {
		matched[i] = string(runes[v[0] : v[0]+v[1]])
	}
	if !reflect.DeepEqual(expectMatched, matched) {
		t.Errorf("\nexpected: %v\nactual: %v", expectMatched, matched)
	}
}

func getKeywords() []string {
	// http://image.gihyo.co.jp/assets/files/book/2010/978-4-7741-4307-1/hugedatabook_samplecode.zip
	// hugedatabook_samplecode/hgdata_example/08/keyword.utf8.uniq.txt
	fp, _ := os.Open("./data/keyword.utf8.uniq.txt")

	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	keywords := make([]string, 216263)
	for i := 0; scanner.Scan(); i++ {
		keywords[i] = scanner.Text()
	}
	return keywords
}
