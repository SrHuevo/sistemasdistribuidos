package main

import (
	"math/rand"
	"strconv"
)

func main() {
	lenPair := 5
	lenList := 1000
	pl := pairList{lenPair, make([]pair, lenList)}
	for i := 0; i < lenList; i++ {
		for j := 0; j < lenPair; j++ {
            if r := rand.Intn(10); r < 5 {
                aux := 
            } 
			if rand.Intn(10) < 4 {
				pl.pair[i] = pair{strconv.Itoa(j), make([]int, lenPair)}
				pl.pair[i].value[j]++
				break
			}
		}
	}
}

type pair struct {
	key   string
	value []int
}

type pairList struct {
	len  int
	pair []pair
}

func (p pairList) Len() int { return len(p.pair) }
func (p pairList) Less(i, j int) bool {
	n1 := 0
	n2 := 0
	for k := 0; k < p.len; k++ {
		n1 += p.pair[i].value[k]
		n2 += p.pair[j].value[k]
	}
	return n1 < n2
}
func (p pairList) Swap(i, j int) { p.pair[i], p.pair[j] = p.pair[j], p.pair[i] }
