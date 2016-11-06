package rendez

import (
	"sync"
)

type cita struct {
	i  interface{}
	wg sync.WaitGroup
}

var rendezMap = make(map[int]*cita)
var mutex = &sync.Mutex{}

func Rendezvous(tag int, val interface{}) interface{} {
	mutex.Lock()
	e, ok := rendezMap[tag]
	if ok {
		rendezMap[tag].wg.Done()
		rendezMap[tag].i = val
		delete(rendezMap, tag)
		mutex.Unlock()
	} else {
		rendezMap[tag].wg.Add(1)
		rendezMap[tag].i = val
		e = rendezMap[tag]
		mutex.Unlock()
		rendezMap[tag].wg.Wait()
	}

	return e.i
}
