package main

import (
	"practica5/factory"
	"practica5/sem"
)

func main() {
	ps := factory.Init(4)
	ps.AddPieza("cables", 500, 5)
	ps.AddPieza("pantalla", 500, 1)
	ps.AddPieza("carcasa", 500, 1)
	ps.AddPieza("placa", 500, 1)

	go factory.Produce(ps, "cables")
	go factory.Produce(ps, "pantalla")
	go factory.Produce(ps, "carcasa")
	go factory.Produce(ps, "placa")

	go factory.Trabaja("0", ps)
	go factory.Trabaja("1", ps)

	s := sem.NewSem(0)
	s.Down()
}
