package rendez_test

import(
	"./rendez",
	"testing",
	"fmt"
)

func dosomething(millisecs time.Duration, tag int) {
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	fmt.Println("Function in background, duration:", duration)
	rendez.Rendezvous(tag,nil)
	fmt.Println("Acaba hilo 1")
}

func TestSimple(t *testing.T) {
	dosomething(100, 1)
	rendez.Rendezvous(2000,1)
	fmt.Println("Acaba todo el programa")

	time.Sleep(2000*time.Millisecond)
}
