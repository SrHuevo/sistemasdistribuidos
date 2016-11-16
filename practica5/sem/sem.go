package sem

import "sync"

type Sem struct {
	m *sync.Mutex
	c *sync.Cond
	ntok int
}

type UpDowner interface {
	NewSem(ntok int) *Sem
	Up()
	Down()
}
