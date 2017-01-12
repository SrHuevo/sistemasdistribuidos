package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sistemasdistribuidos/practica-final/logicclock"
	"sort"
)

type listEvents struct {
	l []logicclock.Event
}

func (le *listEvents) Len() int {
	return len(le.l)
}

func (le *listEvents) Less(i, j int) bool {
	var mod1 float64 = 0
	for _, v := range le.l[i].MapEvents {
		mod1 += math.Pow(float64(v), 2)
	}
	mod1 = math.Sqrt(mod1)
	var mod2 float64 = 0
	for _, v := range le.l[j].MapEvents {
		mod2 += math.Pow(float64(v), 2)
	}
	mod2 = math.Sqrt(mod2)
	return mod1 < mod2
}

func (le *listEvents) Swap(i, j int) {
	le.l[i], le.l[j] = le.l[j], le.l[i]
}

func main() {
	events := new(listEvents)
	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			panic("No se ha podido abrir el fichero " + file)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			e, err := logicclock.FromJson(scanner.Text())

			if err != nil {
				panic("Ha habido un fallo mientras se leía el fichero " + file)
			}
			events.l = append(events.l, e)
		}
	}
	sort.Sort(events)
	for _, event := range events.l {
		fmt.Print(event.Id)
		fmt.Print(": ")
		fmt.Println(event.Log)
	}
}
