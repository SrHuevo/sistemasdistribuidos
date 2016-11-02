package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var words = make(map[string]int)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func saveWordsFiles(v string) {
	f, err := os.Open(v)
	check(err)

	b1 := make([]byte, 1024)
	leidos, err := f.Read(b1)
	check(err)
	for _, v := range strings.Fields(string(b1[:leidos])) {
		words[v]++
	}
}

func main() {
	for _, v := range os.Args[1:] {
		saveWordsFiles(v)
	}

	keys := make([]string, len(words))
	i := 0
	for k := range words {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for _, v := range keys {
		fmt.Printf("%s %d\n\r", v, words[v])
	}
}
