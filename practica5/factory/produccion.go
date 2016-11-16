package factory

func Produce(ps *Piezas, name string) {
	for id := 0; ; id++{
		ps.p[ps.trans[name]].produce(id)
	}
}
