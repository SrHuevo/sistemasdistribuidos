package sem

import (
	"fmt"
	"testing"
	"sync"
	"time"
)

var s *Sem = NewSem(0)

var wg *sync.WaitGroup

func consume(c int) {
	for i := 0; i < 100; i++ {
		s.Down()
		fmt.Println(c)
	}
	wg.Done()
}

func TestSimple(t *testing.T) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	go consume(1)
	wg.Add(1)
	go consume(2)
	time.Sleep(100 * time.Millisecond)
	for i:= 0; i < 200; i++ {
		s.Up();
	}
	wg.Wait()
}
