package main

import (
	"fmt"
	"time"
	"sync"
)

func sEspera(r chan int) {
	for ncnew := 1; ; ncnew++{
		select {
		case r <- ncnew:
			fmt.Println("\tCliente ", ncnew, ": me siento en la sala de espera")
		default:
			fmt.Println("\t\t\tCliente ", ncnew, ": me voy de la barberia, esta llena")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func recep(r, b1, b2 chan int) {
	for {
		c:= <-r
		select {
		case b1 <- c:
		case b2 <- c:
		}
	}
}

func sCorte(nb int, b chan int) {
	for {
		fmt.Println("Barbero ", nb, ": me duermo esperando clientes")
		nc:= <-b
		fmt.Println("\t\tCliente ", nc, ": me corto el pelo")
		fmt.Println("Barbero ", nb, ": empiezo a cortar el pelo")
		time.Sleep(5000 * time.Millisecond)
		fmt.Println("Cliente ", nc, ": termino de cortarme el pelo")
		fmt.Println("Barbero ", nb, ": termino de cortar el pelo")
	}
}

func main() {
	b1 := make(chan int)
	b2 := make(chan int)
	r := make(chan int, 4)//Uno va a estar en manos del recepcionista

	go sCorte(1, b1)
	go sCorte(2, b2)
	time.Sleep(1000 * time.Millisecond)

	go recep(r, b1, b2)
	time.Sleep(1000 * time.Millisecond)

	go sEspera(r)
	time.Sleep(1000 * time.Millisecond)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
