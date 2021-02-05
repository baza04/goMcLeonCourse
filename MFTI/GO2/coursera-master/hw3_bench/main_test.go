package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// запускаем перед основными функциями по разу чтобы файл остался в памяти в файловом кеше
// ioutil.Discard - это ioutil.Writer который никуда не пишет
func init() {
	SlowSearch(ioutil.Discard)
	FastSearch(ioutil.Discard)
}

// -----
// go test -v

func TestSearch(t *testing.T) {
	slowOut := new(bytes.Buffer)
	SlowSearch(slowOut)
	slowResult := slowOut.String()

	fastOut := new(bytes.Buffer)
	FastSearch(fastOut)
	fastResult := fastOut.String()

	if slowResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, slowResult)
	}
}

// -----
// go test -bench . -benchmem

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowSearch(ioutil.Discard)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSearch(ioutil.Discard)
	}
}

var s string

// func BenchmarkConvert(b *testing.B) {
// 	var buf bytes.Buffer
// 	var array = []byte{'m', 'a', 'r', 'k', 'o', 0}
// 	for i := 0; i < b.N; i++ {
// 		buf.Reset()
// 		s = string(array)
// 		buf.WriteString(s)
// 	}
// }
//
// func BenchmarkConvertBad(b *testing.B) {
// 	var buf bytes.Buffer
// 	var array = []byte{'m', 'a', 'r', 'k', 'o'}
// 	for i := 0; i < b.N; i++ {
// 		buf.Reset()
// 		buf.WriteString(string(array))
// 	}
// }