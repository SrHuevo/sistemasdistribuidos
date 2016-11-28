package main

import (
	"fmt"
	"sync"
	"time"
)

type recepType struct {
	salaespera []chan bool
	total      int
	indexIn    int
	indexOut   int
}

const MAX_CLIENTS = 5

func cliente(r chan chan bool, nc int) {
	c := make(chan bool)
	r <- c
	if <-c {
		fmt.Printf("Cliente %d: me siento en la sala de espera\n\r", nc)
		<-c
		fmt.Printf("Cliente %d: me corto el pelo\n\r", nc)
		<-c
		fmt.Printf("Cliente %d: termino de cortarme el pelo\n\r", nc)
	} else {
		fmt.Printf("Cliente %d: me voy de la barberia, esta llena\n\r", nc)
	}
}

func clientes(r chan chan bool) {
	for nc := 1; ; nc++ {
		go cliente(r, nc)

		time.Sleep(1000 * time.Millisecond)
	}
}

func recep(r, b chan chan bool) {
	rtype := &recepType{}
	rtype.salaespera = make([]chan bool, 5)
	rtype.total = 0
	rtype.indexIn = 0
	rtype.indexOut = 0
	for {
		if rtype.total == 0 {
			c := <-r
			enviarASala(c, rtype)
		} else {
			atiendeClientes(rtype, r, b)
		}
	}
}

func atiendeClientes(rtype *recepType, r, b chan chan bool) {
	select {
	case b <- rtype.salaespera[rtype.indexOut]:
		enviarABarbero(rtype)
	case c := <-r:
		if rtype.total == MAX_CLIENTS {
			enviarACalle(c)
		} else {
			enviarASala(c, rtype)
		}
	}
}

func enviarABarbero(rtype *recepType) {
	rtype.indexOut = (rtype.indexOut + 1) % MAX_CLIENTS
	rtype.total--
}

func enviarASala(c chan bool, rtype *recepType) {
	c <- true
	rtype.salaespera[rtype.indexIn] = c
	rtype.indexIn = (rtype.indexIn + 1) % MAX_CLIENTS
	rtype.total++
}

func enviarACalle(c chan bool) {
	c <- false
}

func barbero(nb int, b chan chan bool) {
	for c := range b {
		c <- true
		fmt.Printf("Barbero %d: empiezo a cortar el pelo\r\n", nb)
		time.Sleep(5000 * time.Millisecond)
		fmt.Printf("Barbero %d: termino de cortar el pelo\r\n", nb)
		c <- true
		time.Sleep(time.Millisecond)
		fmt.Printf("Barbero %d: me duermo esperando clientes\r\n", nb)
	}
}

func main() {
	b := make(chan chan bool)
	r := make(chan chan bool)

	go barbero(1, b)
	go barbero(2, b)
	time.Sleep(100 * time.Millisecond)

	go recep(r, b)
	time.Sleep(100 * time.Millisecond)

	go clientes(r)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
