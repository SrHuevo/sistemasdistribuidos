package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	defer f.Close()

	r := bufio.NewReader(f)
	scan := bufio.NewScanner(r)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
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
