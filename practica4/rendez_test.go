package rendez

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg = &sync.WaitGroup{}

func dosomething(millisecs time.Duration, tag int, debug string) {
	fmt.Printf("Empieza: %s\r\n", debug)
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	Rendezvous(tag, nil)
	fmt.Printf("Acaba: %s\r\n", debug)
	wg.Done()
}

func TestSimple(t *testing.T) {
	wg.Add(1)
	go dosomething(100, 1, "primero con tag 1")
	wg.Add(1)
	go dosomething(2000, 2, "primero con tag 2")
	wg.Add(1)
	go dosomething(300, 1, "segundo con tag 1")
	wg.Add(1)
	go dosomething(200, 3, "primero con tag 3")
	wg.Add(1)
	go dosomething(100, 2, "segundo con tag 2")
	wg.Add(1)
	go dosomething(400, 3, "segundo con tag 3")

	wg.Wait()
	fmt.Println("Acaba todo el programa")
}
