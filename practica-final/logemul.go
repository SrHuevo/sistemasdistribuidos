package main

import (
	"fmt"
)

var chans [5]chan int

func run(log string) {
	for i:= 0;; i++ {
		fmt.Println(i, " ", log)
	}
}

func upProcess(log string) {
	chans.append(make(chan int))
	go run(log)
}

func main() {
	upProcess("soy el proceso 1")
	upProcess("soy el proceso 2")
	upProcess("soy el proceso 3")
	upProcess("soy el proceso 4")
	upProcess("soy el proceso 5")
	upProcess("soy el proceso 6")
}
