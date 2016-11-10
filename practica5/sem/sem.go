package sem

import "sync"

var mutex = &sync.Mutex{}

func NewSem(ntok int) *Sem {
	var m sync.Mutex
	c := sync.NewCond(&m)
	sem := Sem{m, c, ntok}
	return &sem
}

func (s *Sem) Up() {
	mutex.Lock()
	s.ntok = s.ntok + 1
	s.c.Signal()
	mutex.Unlock()
}

func (s *Sem) Down() {
	s.m.Lock()
	for s.ntok == 0{
		s.c.Wait()
	}
	s.ntok = s.ntok - 1
	s.m.Unlock()
}
