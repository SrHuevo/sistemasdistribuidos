package factory

import (
	"practica5/sem"
	"strconv"
)

type pieza struct {
	holesem    *sem.Sem
	piezasem   *sem.Sem
	ids        []int
	iConsumido int
	iProducido int
	max int
}

var index int = 0

type Piezas struct {
	names []string
	trans map[string]int
	cant  []int
	p     []*pieza
	total int
}

func Init(tam int) *Piezas {
	ps := Piezas{}
	ps.names = make([]string, tam)
	ps.cant = make([]int, tam)
	ps.trans = make(map[string]int)
	ps.p = make([]*pieza, tam)
	return &ps
}

func (p *pieza) init(max int) {
	p.holesem = sem.NewSem(max)
	p.piezasem = sem.NewSem(0)
	p.ids = make([]int, max)
	p.iConsumido = 0
	p.iProducido = 0
	p.max = max
}

func (ps *Piezas) AddPieza(name string, max int, cant int) {
	ps.names[index] = name
	ps.trans[name] = index
	ps.cant[index] = cant
	p := pieza{}
	p.init(max)
	ps.p[index] = &p
	ps.total += cant
	index++
}

func (p *pieza) produce(id int) {
	p.holesem.Down()
	p.ids[p.iProducido] = id
	p.iProducido = (p.iProducido + 1) % p.max
	p.piezasem.Up()
}

func (p *pieza) get() string {
	p.piezasem.Down()
	id := p.ids[p.iConsumido]
	p.iConsumido = (p.iConsumido + 1) % p.max
	p.holesem.Up()
	return strconv.Itoa(id)
}
